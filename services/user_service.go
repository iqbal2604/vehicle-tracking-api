package services

import (
	"errors"

	"github.com/iqbal2604/vehicle-tracking-api/models"
	"github.com/iqbal2604/vehicle-tracking-api/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(name, email, password string) error {

	existingUser, _ := s.repo.FindByEmail(email)
	if existingUser != nil {
		return errors.New("Email Already Registered")
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	return s.repo.Create(&user)
}
