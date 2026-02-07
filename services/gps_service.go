package services

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/iqbal2604/vehicle-tracking-api/dtos"
	"github.com/iqbal2604/vehicle-tracking-api/models"
	"github.com/iqbal2604/vehicle-tracking-api/repositories"
	websocketpkg "github.com/iqbal2604/vehicle-tracking-api/websocket"
)

type GPSService struct {
	gpsRepo     *repositories.GPSRepository
	vehicleRepo *repositories.VehicleRepository
	hub         *websocketpkg.Hub
}

func NewGPSService(gpsRepo *repositories.GPSRepository, vehicleRepo *repositories.VehicleRepository, hub *websocketpkg.Hub) *GPSService {
	return &GPSService{
		gpsRepo:     gpsRepo,
		vehicleRepo: vehicleRepo,
		hub:         hub,
	}
}

func (s *GPSService) CreateLocation(userID uint, loc *models.GPSLocation) error {

	_, err := s.vehicleRepo.FindByID(userID, loc.VehicleID)
	if err != nil {
		return errors.New("Vehicle Not Found")
	}
	if err := s.gpsRepo.Create(loc); err != nil {
		return err
	}

	data, _ := json.Marshal(dtos.ToGPSResponse(*loc))
	s.hub.Broadcast <- websocketpkg.WSMessage{
		UserID:    userID,
		VehicleID: loc.VehicleID,
		Data:      data,
	}

	return nil

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

func (s *GPSService) GetVehicleStatus(userID, vehicleID uint) (string, error) {
	loc, err := s.gpsRepo.GetLastByVehicleID(vehicleID)
	if err != nil {
		return "offline", nil
	}
	if time.Since(loc.CreatedAt) <= 30*time.Second {
		return "online", nil
	}
	return "offline", nil
}
