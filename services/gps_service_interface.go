package services

import models "github.com/iqbal2604/vehicle-tracking-api/models/domain"

type GPSService interface {
	CreateLocation(userID uint, loc *models.GPSLocation) error
	GetLastLocation(userID uint, vehicleID uint) (*models.GPSLocation, error)
	GetHistory(userID uint, vehicleID uint) ([]models.GPSLocation, error)
	GetVehicleStatus(userID, vehicleID uint) (string, error)
}
