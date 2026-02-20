package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/iqbal2604/vehicle-tracking-api/models"
	"github.com/iqbal2604/vehicle-tracking-api/repositories"
	"github.com/redis/go-redis/v9"
)

type VehicleServiceImpl struct {
	repo *repositories.VehicleRepository
	rdb  *redis.Client
}

func NewVehicleService(repo *repositories.VehicleRepository, rdb *redis.Client) VehicleService {
	return &VehicleServiceImpl{repo: repo, rdb: rdb}
}

func (s *VehicleServiceImpl) CreateVehicle(userID uint, v *models.Vehicle) error {
	v.UserID = userID
	err := s.repo.Create(v)
	if err == nil {
		s.rdb.Del(context.Background(), fmt.Sprintf("vehicles:user:%d", userID))
	}
	return err
}

func (s *VehicleServiceImpl) GetVehicleByID(userID uint, id uint) (*models.Vehicle, error) {
	return s.repo.FindByID(userID, id)
}

func (s *VehicleServiceImpl) ListVehiclesByUser(userID uint) ([]models.Vehicle, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("vehicles:user:%d", userID)

	//Ambil data dari redis
	cachedData, err := s.rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		var vehicles []models.Vehicle
		json.Unmarshal([]byte(cachedData), &vehicles)
		return vehicles, nil
	}

	//Kalo ga ada di redis ambil dari db
	vehicles, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	//SImpan Cache ke redis 5 menit
	jsonData, _ := json.Marshal(vehicles)
	s.rdb.Set(ctx, cacheKey, jsonData, 5*time.Minute)

	return vehicles, nil
}

func (s *VehicleServiceImpl) UpdateVehicle(userID uint, v *models.Vehicle) error {
	existing, err := s.repo.FindByID(userID, v.ID)
	if err != nil {
		return err
	}

	if existing.UserID != userID {
		return errors.New("Unauthorized")
	}
	s.rdb.Del(context.Background(), fmt.Sprintf("vehicles:user:%d", userID))
	v.UserID = userID
	return s.repo.Update(v)
}

func (s *VehicleServiceImpl) DeleteVehicle(userID, id uint) error {
	existing, err := s.repo.FindByID(userID, id)
	if err != nil {
		return err
	}

	if existing.UserID != userID {
		return errors.New("Unauthorized")
	}

	s.rdb.Del(context.Background(), fmt.Sprintf("vehicles:user:%d", userID))

	return s.repo.Delete(userID, id)
}

func (s *VehicleServiceImpl) ListAllVehicles() ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	if err := s.repo.DB.Preload("User").Find(&vehicles).Error; err != nil {
		return nil, err
	}
	return vehicles, nil
}
