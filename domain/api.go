package domain

// cihazın bağlandığı dış API
type SensorAPI struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	DeviceID string `gorm:"not null"` // FK
	Device   SensorDevice
	Endpoint string `gorm:"size:255;not null"` // Örn: "http://localhost:8081/temp"
	Method   string `gorm:"size:16;not null"`  // GET / POST
}
