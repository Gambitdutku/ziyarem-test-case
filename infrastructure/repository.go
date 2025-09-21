package infrastructure

import (
	"context"
	"ziyaremtestcase/domain"

	"gorm.io/gorm"
)

// DB adaptörr
type SensorRepository struct {
	db *gorm.DB
}

func NewSensorRepository(db *gorm.DB) *SensorRepository {
	return &SensorRepository{db: db}
}

func (r *SensorRepository) Save(ctx context.Context, data *domain.SensorData) error {
	return r.db.WithContext(ctx).Save(data).Error
}

func (r *SensorRepository) FindByID(ctx context.Context, id string) (*domain.SensorData, error) {
	var d domain.SensorData
	if err := r.db.WithContext(ctx).First(&d, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &d, nil
}
