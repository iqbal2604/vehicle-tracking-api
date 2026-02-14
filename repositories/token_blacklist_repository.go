package repositories

import (
	"time"

	"github.com/iqbal2604/vehicle-tracking-api/models"
	"gorm.io/gorm"
)

type TokenBlacklistRepository struct {
	DB *gorm.DB
}

func NewTokenBlacklistRepository(db *gorm.DB) *TokenBlacklistRepository {
	return &TokenBlacklistRepository{
		DB: db,
	}
}

func (r *TokenBlacklistRepository) AddToken(token string, expiresAt time.Time) error {
	blacklist := models.TokenBlacklist{
		Token:     token,
		ExpiresAt: expiresAt,
	}
	return r.DB.Create(&blacklist).Error
}

func (r *TokenBlacklistRepository) IsTokenBlacklisted(token string) (bool, error) {
	var count int64
	err := r.DB.Model(&models.TokenBlacklist{}).Where("token = ?", token).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
