package repositories

import (
	"github.com/iqbal2604/vehicle-tracking-api/models"
	"gorm.io/gorm"
)

type VehicleRepository struct {
	DB *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) *VehicleRepository {
	return &VehicleRepository{DB: db}
}

func (r *VehicleRepository) Create(vehicle *models.Vehicle) error {
	if err := r.DB.Create(vehicle).Error; err != nil {
		return err
	}
	return r.DB.Preload("User").Find(vehicle).Error
}

func (r *VehicleRepository) FindByID(userID uint, id uint) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	err := r.DB.Preload("User").Where("id = ? AND user_id = ?", id, userID).First(&vehicle).Error
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (r *VehicleRepository) FindByUserID(userID uint) ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	err := r.DB.Preload("User").Where("user_id = ?", userID).Find(&vehicles).Error
	if err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (r *VehicleRepository) Update(vehicle *models.Vehicle) error {
	return r.DB.Save(vehicle).Error
}

func (r *VehicleRepository) Delete(userID uint, id uint) error {
	return r.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Vehicle{}, id).Error
}
