package domain

import "time"

// SensorData -> ölçüm sonuçları
type SensorData struct {
	ID        string `gorm:"primaryKey;size:64"`
	DeviceID  string `gorm:"size:64;not null"` // Foreign key -> SensorDevice.DeviceID
	Device    SensorDevice
	Value     float64   `gorm:"not null"`
	Timestamp time.Time `gorm:"not null"`
}

// Sensör interface’i (her tip implemente eder)
type Sensor interface {
	Read(id string) (*SensorData, error)
	Type() string
}
