package models

import "time"

type Geofence struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Radius    float64   `json:"radius"`
	Type      string    `json:"type" gorm:"default:safe_zone"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"user"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
