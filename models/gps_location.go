package models

type GPSLocation struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Speed     float64 `json:"speed"`

	VehicleID uint `json:"vehicle_id"`
	Vehicle   Vehicle

	CreatedAt int64 `gorm:"autoCreateTime"`
}
