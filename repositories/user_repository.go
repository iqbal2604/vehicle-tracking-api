package repositories

import (
	"github.com/iqbal2604/vehicle-tracking-api/config"
	"github.com/iqbal2604/vehicle-tracking-api/models"
)

type UserRepository struct{}

func (r *UserRepository) Create(user *models.User) error {
	return config.DB.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
