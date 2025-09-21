package application

import (
	"context"
	"fmt"
	"time"
	"ziyaremtestcase/domain"

	"github.com/sirupsen/logrus"
)

type Cache interface {
	Get(ctx context.Context, key string) (*domain.SensorData, error)
	Set(ctx context.Context, key string, data *domain.SensorData, ttl time.Duration) error
}

type Repository interface {
	Save(ctx context.Context, data *domain.SensorData) error
	FindByID(ctx context.Context, id string) (*domain.SensorData, error)
}

type AppService struct {
	cache    Cache
	repo     Repository
	cb       *CircuitBreaker
	logger   *logrus.Logger
	cacheTTL time.Duration
	retries  int
}

func NewAppService(cache Cache, repo Repository, cb *CircuitBreaker, logger *logrus.Logger) *AppService {
	return &AppService{
		cache:    cache,
		repo:     repo,
		cb:       cb,
		logger:   logger,
		cacheTTL: 10 * time.Second,
		retries:  3,
	}
}

func (a *AppService) GetSensorData(ctx context.Context, sensor domain.Sensor, id string) (*domain.SensorData, error) {
	key := fmt.Sprintf("%s:%s", sensor.Type(), id)

	// 1. Cache kontrolü
	if d, _ := a.cache.Get(ctx, key); d != nil {
		a.logger.Infof("Cache hit: %s", key)
		return d, nil
	}

	// 2. Circuit breaker kontrolü
	if !a.cb.Allow(sensor.Type()) {
		a.logger.Warnf("Circuit breaker açık: %s", sensor.Type())
		return nil, fmt.Errorf("circuit breaker open for %s", sensor.Type())
	}

	var data *domain.SensorData
	var err error

	// 3. yeniden dene
	for attempt := 1; attempt <= a.retries; attempt++ {
		a.logger.Infof("Attempt %d to read sensor %s (%s)", attempt, sensor.Type(), id)
		data, err = sensor.Read(id)
		if err == nil {
			// Başarılı
			a.cb.Success(sensor.Type())
			if saveErr := a.repo.Save(ctx, data); saveErr != nil {
				a.logger.Errorf("DB save failed: %v", saveErr)
			}
			if cacheErr := a.cache.Set(ctx, key, data, a.cacheTTL); cacheErr != nil {
				a.logger.Warnf("Cache set failed: %v", cacheErr)
			}
			return data, nil
		}
		a.logger.Warnf("Sensor read error: %v", err)
		time.Sleep(500 * time.Millisecond)
	}

	// 4. Başarısız
	a.cb.Failure(sensor.Type())
	a.logger.Errorf("All retries failed for sensor %s (%s)", sensor.Type(), id)
	return nil, err
}
