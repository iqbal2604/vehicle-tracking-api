package dtos

import "github.com/iqbal2604/vehicle-tracking-api/models"

type GPSResponse struct {
	ID        uint    `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Speed     float64 `json:"speed"`
	CreatedAt int64   `json:"created_at"`
	VehicleID uint    `json:"vehicle_id"`
}

type GPSHistoryResponse struct {
	VehicleID uint          `json:"vehicle_id"`
	Locations []GPSResponse `json:"locations"`
}

func ToGPSResponse(g models.GPSLocation) GPSResponse {
	return GPSResponse{
		ID:        g.ID,
		Latitude:  g.Latitude,
		Longitude: g.Longitude,
		Speed:     g.Speed,
		CreatedAt: g.CreatedAt,
		VehicleID: g.VehicleID,
	}
}
