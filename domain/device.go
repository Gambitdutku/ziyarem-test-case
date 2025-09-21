package domain

// fiziksel cihazlar
type SensorDevice struct {
	DeviceID string `gorm:"primaryKey;size:64"` // Örn: "temp-001"
	TypeID   uint   `gorm:"not null"`           //fk
	Type     SensorType
	Location string `gorm:"size:128"` // bulunduğu yer
}
