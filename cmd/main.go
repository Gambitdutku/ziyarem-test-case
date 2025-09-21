package main

import (
	"log"
	"ziyaremtestcase/domain"
	"ziyaremtestcase/infrastructure"
)

func main() {
	db := infrastructure.NewDB()

	// Tipler
	types := []domain.SensorType{
		{Name: "temperature"},
		{Name: "humidity"},
		{Name: "airquality"},
	}
	for _, t := range types {
		db.FirstOrCreate(&t, domain.SensorType{Name: t.Name})
	}

	// Cihazlar
	devices := []domain.SensorDevice{
		{DeviceID: "temp-001", TypeID: 1, Location: "Room A"},
		{DeviceID: "hum-001", TypeID: 2, Location: "Room B"},
		{DeviceID: "air-001", TypeID: 3, Location: "Room C"},
	}
	for _, d := range devices {
		db.FirstOrCreate(&d, domain.SensorDevice{DeviceID: d.DeviceID})
	}

	// API tanımları
	apis := []domain.SensorAPI{
		{DeviceID: "temp-001", Endpoint: "http://localhost:8081/temp", Method: "GET"},
		{DeviceID: "hum-001", Endpoint: "http://localhost:8082/hum", Method: "GET"},
		{DeviceID: "air-001", Endpoint: "http://localhost:8083/air", Method: "GET"},
	}
	for _, a := range apis {
		db.FirstOrCreate(&a, domain.SensorAPI{DeviceID: a.DeviceID})
	}

	// Kontrol
	var devicesCheck []domain.SensorDevice
	db.Preload("Type").Find(&devicesCheck)
	for _, d := range devicesCheck {
		log.Printf("📡 Device: %s Type=%s Location=%s", d.DeviceID, d.Type.Name, d.Location)
	}

	// Kontrol
	var apisCheck []domain.SensorAPI
	db.Preload("Device").Find(&apisCheck)
	for _, a := range apisCheck {
		log.Printf("🌐 API for %s -> %s %s", a.Device.DeviceID, a.Method, a.Endpoint)
	}
}
