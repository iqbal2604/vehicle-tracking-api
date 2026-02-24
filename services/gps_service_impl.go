package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/iqbal2604/vehicle-tracking-api/dtos"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
	"github.com/iqbal2604/vehicle-tracking-api/models"
	"github.com/iqbal2604/vehicle-tracking-api/repositories"
	websocketpkg "github.com/iqbal2604/vehicle-tracking-api/websocket"
	"github.com/redis/go-redis/v9"
)

type GPSServiceImpl struct {
	gpsRepo      *repositories.GPSRepository
	vehicleRepo  *repositories.VehicleRepository
	geofenceRepo *repositories.GeofenceRepository
	hub          *websocketpkg.Hub
	rdb          *redis.Client
}

func NewGPSService(gpsRepo *repositories.GPSRepository, vehicleRepo *repositories.VehicleRepository, geofenceRepo *repositories.GeofenceRepository, hub *websocketpkg.Hub, rdb *redis.Client) GPSService {
	return &GPSServiceImpl{
		gpsRepo:      gpsRepo,
		vehicleRepo:  vehicleRepo,
		geofenceRepo: geofenceRepo,
		hub:          hub,
		rdb:          rdb,
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
	go s.checkgeoFences(userID, loc)

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
	ctx := context.Background()
	cacheKey := fmt.Sprintf("vehicle:last:%d", vehicleID)
	var lastloc models.GPSLocation

	cacheData, err := s.rdb.Get(ctx, cacheKey).Result()
	if err != nil {
		json.Unmarshal([]byte(cacheData), &lastloc)
	} else {
		loc, err := s.gpsRepo.GetLastByVehicleID(vehicleID)
		if err != nil {
			return "offline", nil
		}
		if loc == nil {
			return "offline", nil
		}
		lastloc = *loc
	}
	if time.Since(lastloc.CreatedAt) <= 30*time.Second {
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

func (s *GPSServiceImpl) CreateGeofence(geofence *models.Geofence) error {
	return s.geofenceRepo.Create(geofence)
}

func (s *GPSServiceImpl) ListGeofences(userID uint) ([]models.Geofence, error) {
	return s.geofenceRepo.FindByUserID(userID)
}

func (s *GPSServiceImpl) DeleteGeofence(id uint, userID uint) error {
	return s.geofenceRepo.Delete(id, userID)
}

func (s *GPSServiceImpl) checkgeoFences(userID uint, loc *models.GPSLocation) {
	geofences, err := s.geofenceRepo.FindByUserID(userID)
	if err != nil {
		return
	}

	for _, gf := range geofences {
		distance := helpers.CalculateDistance(loc.Latitude, loc.Longitude, gf.Latitude, gf.Longitude)

		isOutside := distance > gf.Radius

		// Logic: Jika ini 'safe_zone' tapi kendaraan di luar (isOutside = true)
		// Atau jika ini 'restricted_area' tapi kendaraan di dalam (isOutside = false)
		violation := false
		message := ""

		if gf.Type == "safe_zone" && isOutside {
			violation = true
			message = fmt.Sprintf("⚠️ ALERT: Vehicle %d KELUAR dari area aman (%s)!", loc.VehicleID, gf.Name)
		} else if gf.Type == "restricted_area" && !isOutside {
			violation = true
			message = fmt.Sprintf("⛔ ALERT: Vehicle %d MASUK ke area terlarang (%s)!", loc.VehicleID, gf.Name)
		}

		if violation {
			// Kirim Alert via WebSocket
			alertData, _ := json.Marshal(map[string]interface{}{
				"type":       "GEOFENCE_ALERT",
				"vehicle_id": loc.VehicleID,
				"message":    message,
				"lat":        loc.Latitude,
				"lng":        loc.Longitude,
			})

			s.hub.Broadcast <- websocketpkg.WSMessage{
				UserID: userID,
				Data:   alertData,
			}
		}
	}
}
