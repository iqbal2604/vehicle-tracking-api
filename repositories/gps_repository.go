package repositories

import (
	"github.com/iqbal2604/vehicle-tracking-api/models"
	"gorm.io/gorm"
)

type GPSRepository struct {
	DB *gorm.DB
}

func NewGPSRepository(db *gorm.DB) *GPSRepository {
	return &GPSRepository{DB: db}
}

// Insert Lokasi Baru
func (r *GPSRepository) Create(location *models.GPSLocation) error {
	return r.DB.Create(location).Error
}

// Ambil Lokasi terakhir Vehicle
func (r *GPSRepository) GetLastByVehicleID(vehicleID uint) (*models.GPSLocation, error) {
	var loc models.GPSLocation

	err := r.DB.Where("vehicle_id = ?", vehicleID).Order("created_at desc").First(&loc).Error
	if err != nil {
		return nil, err
	}
	return &loc, nil
}

// Ambil History Lokasi
func (r *GPSRepository) GetHistory(vehicleID uint) ([]models.GPSLocation, error) {
	var locations []models.GPSLocation

	err := r.DB.Where("vehicle_id = ?", vehicleID).Order("created_at asc").Find(&locations).Error

	if err != nil {
		return nil, err
	}

	return locations, nil
}
