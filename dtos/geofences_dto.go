package dtos

import "github.com/iqbal2604/vehicle-tracking-api/models"

type GeofenceRequest struct {
	Name      string  `json:"name" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
	Radius    float64 `json:"radius" validate:"required,min=10"`
	Type      string  `json:"type" validate:"required,onof=safe_zone restricted_area"`
}

type GeofenceResponse struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius    float64 `json:"radius"`
	Type      string  `json:"type"`
}

func ToGeofenceResponse(g models.Geofence) GeofenceResponse {
	return GeofenceResponse{
		ID:        g.ID,
		Name:      g.Name,
		Latitude:  g.Latitude,
		Longitude: g.Longitude,
		Radius:    g.Radius,
		Type:      g.Type,
	}
}

func ToGeofenceListResponse(geofences []models.Geofence) []GeofenceResponse {
	var responses []GeofenceResponse
	for _, g := range geofences {
		responses = append(responses, ToGeofenceResponse(g))
	}
	return responses
}
