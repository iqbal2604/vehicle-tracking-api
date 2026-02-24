package repositories

import (
	"github.com/iqbal2604/vehicle-tracking-api/models"
	"gorm.io/gorm"
)

type GeofenceRepository struct {
	db *gorm.DB
}

func NewGeofenceRepository(db *gorm.DB) *GeofenceRepository {
	return &GeofenceRepository{db: db}
}

func (r *GeofenceRepository) Create(geofence *models.Geofence) error {
	return r.db.Create(geofence).Error
}

func (r *GeofenceRepository) FindByUserID(userID uint) ([]models.Geofence, error) {
	var geofences []models.Geofence
	err := r.db.Where("user_id = ?", userID).Find(&geofences).Error
	return geofences, err
}

func (r *GeofenceRepository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ? ", id, userID).Delete(&models.Geofence{}).Error
}
