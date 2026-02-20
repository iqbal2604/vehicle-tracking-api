package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/iqbal2604/vehicle-tracking-api/dtos"
	"github.com/iqbal2604/vehicle-tracking-api/models"
	"github.com/iqbal2604/vehicle-tracking-api/repositories"
	websocketpkg "github.com/iqbal2604/vehicle-tracking-api/websocket"
	"github.com/redis/go-redis/v9"
)

type GPSServiceImpl struct {
	gpsRepo     *repositories.GPSRepository
	vehicleRepo *repositories.VehicleRepository
	hub         *websocketpkg.Hub
	rdb         *redis.Client
}

func NewGPSService(gpsRepo *repositories.GPSRepository, vehicleRepo *repositories.VehicleRepository, hub *websocketpkg.Hub, rdb *redis.Client) GPSService {
	return &GPSServiceImpl{
		gpsRepo:     gpsRepo,
		vehicleRepo: vehicleRepo,
		hub:         hub,
		rdb:         rdb,
	}
}

func (s *GPSServiceImpl) CreateLocation(userID uint, loc *models.GPSLocation) error {

	_, err := s.vehicleRepo.FindByID(userID, loc.VehicleID)
	if err != nil {
		return errors.New("Vehicle Not Found")
	}

	if err := s.gpsRepo.Create(loc); err != nil {
		return err
	}

	//Simpan Lokasi Terakhir ke Redis
	ctx := context.Background()
	cacheKey := fmt.Sprintf("vehicle:last:%d", loc.VehicleID)
	locData, _ := json.Marshal(loc)
	s.rdb.Set(ctx, cacheKey, locData, 24*time.Hour)

	data, _ := json.Marshal(dtos.ToGPSResponse(*loc))
	s.hub.Broadcast <- websocketpkg.WSMessage{
		UserID:    userID,
		VehicleID: loc.VehicleID,
		Data:      data,
	}

	return nil

}

func (s *GPSServiceImpl) GetLastLocation(userID uint, vehicleID uint) (*models.GPSLocation, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("vehicle:last:%d", vehicleID)

	//Cek Redis Terlebih Dahulu
	cacheData, err := s.rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		var loc models.GPSLocation
		json.Unmarshal([]byte(cacheData), &loc)
		return &loc, nil
	}
	return s.gpsRepo.GetLastByVehicleID(vehicleID)
}

func (s *GPSServiceImpl) GetHistory(userID uint, vehicleID uint) ([]models.GPSLocation, error) {
	_, err := s.vehicleRepo.FindByID(userID, vehicleID)
	if err != nil {
		return nil, errors.New("Vehicle Not Found")
	}

	return s.gpsRepo.GetHistory(vehicleID)
}

func (s *GPSServiceImpl) GetVehicleStatus(userID, vehicleID uint) (string, error) {
	loc, err := s.gpsRepo.GetLastByVehicleID(vehicleID)
	if err != nil {
		return "offline", nil
	}
	if time.Since(loc.CreatedAt) <= 30*time.Second {
		return "online", nil
	}
	return "offline", nil
}

func (s *GPSServiceImpl) GetLastLocationAdmin(vehicleID uint) (*models.GPSLocation, error) {
	return s.gpsRepo.GetLastByVehicleID(vehicleID)
}

func (s *GPSServiceImpl) GetHistoryAdmin(vehicleID uint) ([]models.GPSLocation, error) {
	return s.gpsRepo.GetHistory(vehicleID)
}
