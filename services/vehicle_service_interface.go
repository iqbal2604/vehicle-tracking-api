package services

import "github.com/iqbal2604/vehicle-tracking-api/models"

type VehicleService interface {
	CreateVehicle(userID uint, v *models.Vehicle) error
	GetVehicleByID(userID uint, id uint) (*models.Vehicle, error)
	ListVehiclesByUser(userID uint) ([]models.Vehicle, error)
	UpdateVehicle(userID uint, v *models.Vehicle) error
	DeleteVehicle(userID, id uint) error
	ListAllVehicles() ([]models.Vehicle, error)
}
