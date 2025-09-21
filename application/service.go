package application

import (
	"context"
	"errors"
	"fmt"
	"time"
	"ziyaremtestcase/domain"
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
	cacheTTL time.Duration
}

func NewAppService(cache Cache, repo Repository) *AppService {
	return &AppService{cache: cache, repo: repo, cacheTTL: 10 * time.Second}
}

func (a *AppService) GetSensorData(ctx context.Context, sensor domain.Sensor, id string) (*domain.SensorData, error) {
	cacheKey := fmt.Sprintf("%s:%s", sensor.Type(), id)

	// 1. Cache kontrolü
	if d, _ := a.cache.Get(ctx, cacheKey); d != nil {
		return d, nil
	}

	// 2. Sensörü oku
	data, err := sensor.Read(id)
	if err != nil {
		return nil, errors.New("sensor read failed: " + err.Error())
	}

	// 3. DB’ye kaydet
	if err := a.repo.Save(ctx, data); err != nil {
		return nil, errors.New("db save failed: " + err.Error())
	}

	// 4. Cache’e yaz
	_ = a.cache.Set(ctx, cacheKey, data, a.cacheTTL)

	return data, nil
}
