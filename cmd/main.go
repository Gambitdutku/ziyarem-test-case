package main

import (
	"context"
	"ziyaremtestcase/application"
	"ziyaremtestcase/domain"
	"ziyaremtestcase/infrastructure"
	"ziyaremtestcase/sensors"

	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	db := infrastructure.NewDB()
	cache := infrastructure.NewRedisCache()
	repo := infrastructure.NewSensorRepository(db)

	cb := application.NewCircuitBreaker(3, 30*time.Second)
	service := application.NewAppService(cache, repo, cb, logger)

	// DB’den sensör API listesi al
	var apis []domain.SensorAPI
	db.Preload("Device.Type").Find(&apis)

	// Her sensörü oku
	for _, api := range apis {
		var sensor domain.Sensor
		switch api.Device.Type.Name {
		case "temperature":
			sensor = &sensors.TempSensor{Endpoint: api.Endpoint}
		case "humidity":
			sensor = &sensors.HumiditySensor{Endpoint: api.Endpoint}
		case "airquality":
			sensor = &sensors.AirQualitySensor{Endpoint: api.Endpoint}
		default:
			logger.Warnf("Desteklenmeyen sensör tipi: %s", api.Device.Type.Name)
			continue
		}

		data, err := service.GetSensorData(context.Background(), sensor, api.DeviceID)
		if err != nil {
			logger.Errorf("Sensor %s failed: %v", api.DeviceID, err)
			continue
		}
		logger.Infof("SensorData okundu ve kaydedildi: %+v", data)
	}
}
