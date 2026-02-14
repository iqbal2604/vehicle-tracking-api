package services

import (
	"github.com/iqbal2604/vehicle-tracking-api/models"
)

type UserService interface {
	Register(name, email, password, role string) (*models.User, error)
	Login(email, password string) (string, uint, error)
	GetProfile(userID uint) (*models.User, error)
	Logout(token string, userID uint) error
}
