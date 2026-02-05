package services

import (
	"errors"

	"github.com/iqbal2604/vehicle-tracking-api/models"
	"github.com/iqbal2604/vehicle-tracking-api/repositories"
)

type GPSService struct {
	gpsRepo     *repositories.GPSRepository
	vehicleRepo *repositories.VehicleRepository
}

func NewGPSService(gpsRepo *repositories.GPSRepository, vehicleRepo *repositories.VehicleRepository) *GPSService {
	return &GPSService{
		gpsRepo:     gpsRepo,
		vehicleRepo: vehicleRepo,
	}
}

func (s *GPSService) CreateLocation(userID uint, loc *models.GPSLocation) error {

	_, err := s.vehicleRepo.FindByID(userID, loc.VehicleID)
	if err != nil {
		return errors.New("Vehicle Not Found")
	}
	return s.gpsRepo.Create(loc)
}

func (s *GPSService) GetLastLocation(userID uint, vehicleID uint) (*models.GPSLocation, error) {
	_, err := s.vehicleRepo.FindByID(userID, vehicleID)

	if err != nil {
		return nil, errors.New("Vehicle Not Found")
	}

	return s.gpsRepo.GetLastByVehicleID(vehicleID)
}

func (s *GPSService) GetHistory(userID uint, vehicleID uint) ([]models.GPSLocation, error) {
	_, err := s.vehicleRepo.FindByID(userID, vehicleID)
	if err != nil {
		return nil, errors.New("Vehicle Not Found")
	}

	return s.gpsRepo.GetHistory(vehicleID)
}
