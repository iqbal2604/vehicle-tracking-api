package dtos

import models "github.com/iqbal2604/vehicle-tracking-api/models/domain"

type VehicleResponse struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Plate  string `json:"plate"`
	Type   string `json:"type"`
	Model  string `json:"model"`
	Color  string `json:"color"`
}

func ToVehicleResponse(v models.Vehicle) VehicleResponse {
	return VehicleResponse{
		ID:     v.ID,
		UserID: v.UserID,
		Name:   v.Name,
		Plate:  v.Plate,
		Type:   v.Type,
		Model:  v.Model,
		Color:  v.Color,
	}
}
