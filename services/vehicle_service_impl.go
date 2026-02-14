package services

import (
	"errors"

	"github.com/iqbal2604/vehicle-tracking-api/models"
	"github.com/iqbal2604/vehicle-tracking-api/repositories"
)

type VehicleServiceImpl struct {
	repo *repositories.VehicleRepository
}

func NewVehicleService(repo *repositories.VehicleRepository) VehicleService {
	return &VehicleServiceImpl{repo: repo}
}

func (s *VehicleServiceImpl) CreateVehicle(userID uint, v *models.Vehicle) error {
	v.UserID = userID
	return s.repo.Create(v)
}

func (s *VehicleServiceImpl) GetVehicleByID(userID uint, id uint) (*models.Vehicle, error) {
	return s.repo.FindByID(userID, id)
}

func (s *VehicleServiceImpl) ListVehiclesByUser(userID uint) ([]models.Vehicle, error) {
	return s.repo.FindByUserID(userID)
}

func (s *VehicleServiceImpl) UpdateVehicle(userID uint, v *models.Vehicle) error {
	existing, err := s.repo.FindByID(userID, v.ID)
	if err != nil {
		return err
	}

	if existing.UserID != userID {
		return errors.New("Unauthorized")
	}
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

	return s.repo.Delete(userID, id)
}

func (s *VehicleServiceImpl) ListAllVehicles() ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	if err := s.repo.DB.Preload("User").Find(&vehicles).Error; err != nil {
		return nil, err
	}
	return vehicles, nil
}
