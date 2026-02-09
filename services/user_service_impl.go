package services

import (
	"errors"

	"github.com/iqbal2604/vehicle-tracking-api/helpers"
	models "github.com/iqbal2604/vehicle-tracking-api/models/domain"
	"github.com/iqbal2604/vehicle-tracking-api/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) UserService {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) Register(name, email, password, role string) error {
	if role == "" {
		role = "driver"
	}

	existingUser, _ := s.repo.FindByEmail(email)
	if existingUser != nil {
		return errors.New("Email Already Registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}

	return s.repo.Create(&user)
}

func (s *UserServiceImpl) Login(email, password string) (string, error) {

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("Email Not Found")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("Wrong Password")
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *UserServiceImpl) GetProfile(userID uint) (*models.User, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
