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
	return r.DB.Create(vehicle).Error
}

func (r *VehicleRepository) FindByID(id uint) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	err := r.DB.First(&vehicle, id).Error
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (r *VehicleRepository) FindByUserID(userID uint) ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	err := r.DB.Where("user_id = ?", userID).Find(&vehicles).Error
	if err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (r *VehicleRepository) Update(vehicle *models.Vehicle) error {
	return r.DB.Save(vehicle).Error
}

func (r *VehicleRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Vehicle{}, id).Error
}
