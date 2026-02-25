package services

import "github.com/iqbal2604/vehicle-tracking-api/models"

type GPSService interface {
	//GPS Tracking
	CreateLocation(userID uint, loc *models.GPSLocation) error
	GetLastLocation(userID uint, vehicleID uint) (*models.GPSLocation, error)
	GetHistory(userID uint, vehicleID uint, start string, end string) ([]models.GPSLocation, error)
	GetVehicleStatus(userID, vehicleID uint) (string, error)

	//Geofence Management
	CreateGeofence(geofence *models.Geofence) error
	ListGeofences(userID uint) ([]models.Geofence, error)
	DeleteGeofence(id uint, userID uint) error

	// Admin methods (bypass ownership check)
	GetLastLocationAdmin(vehicleID uint) (*models.GPSLocation, error)
	GetHistoryAdmin(vehicleID uint, start string, end string) ([]models.GPSLocation, error)
}
