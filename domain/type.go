package domain

// Sensör tipleri
type SensorType struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"unique;size:64;not null"` // Örnek: "temperature"
}
