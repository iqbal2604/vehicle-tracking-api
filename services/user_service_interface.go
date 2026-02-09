package services

import models "github.com/iqbal2604/vehicle-tracking-api/models/domain"

type UserService interface {
	Register(name, email, password, role string) error
	Login(email, password string) (string, error)
	GetProfile(userID uint) (*models.User, error)
}
