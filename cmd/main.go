package main

import (
	"context"
	"log"
	"ziyaremtestcase/application"
	"ziyaremtestcase/domain"
	"ziyaremtestcase/infrastructure"
	"ziyaremtestcase/sensors"
)

func main() {
	db := infrastructure.NewDB()
	cache := infrastructure.NewRedisCache()
	repo := infrastructure.NewSensorRepository(db)

	service := application.NewAppService(cache, repo)

	// Sensörleri veri tabanından çek
	var apis []domain.SensorAPI
	db.Preload("Device.Type").Find(&apis)

	// Sensör tipine göre adaptörr seç
	for _, api := range apis {
		var sensor domain.Sensor
		switch api.Device.Type.Name {
		case "temperature":
			sensor = &sensors.TempSensor{Endpoint: api.Endpoint}
		case "humidity":
			sensor = &sensors.HumiditySensor{Endpoint: api.Endpoint}
		default:
			log.Printf("desteklenmeyen sensör tipi: %s", api.Device.Type.Name)
			continue
		}

		// Sensör verisini oku
		data, err := service.GetSensorData(context.Background(), sensor, api.DeviceID)
		if err != nil {
			log.Printf("Sensör (%s) verisi alınamadı: %v", api.DeviceID, err)
			continue
		}

		log.Printf("SensorData kaydedildi: %+v", data)
	}
}
