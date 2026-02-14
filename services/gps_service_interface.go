package services

import "github.com/iqbal2604/vehicle-tracking-api/models"

type GPSService interface {
	CreateLocation(userID uint, loc *models.GPSLocation) error
	GetLastLocation(userID uint, vehicleID uint) (*models.GPSLocation, error)
	GetHistory(userID uint, vehicleID uint) ([]models.GPSLocation, error)
	GetVehicleStatus(userID, vehicleID uint) (string, error)

	// Admin methods (bypass ownership check)
	GetLastLocationAdmin(vehicleID uint) (*models.GPSLocation, error)
	GetHistoryAdmin(vehicleID uint) ([]models.GPSLocation, error)
}
