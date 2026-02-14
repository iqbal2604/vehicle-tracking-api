package services

import (
	"errors"
	"time"

	"github.com/iqbal2604/vehicle-tracking-api/helpers"
	"github.com/iqbal2604/vehicle-tracking-api/models"
	"github.com/iqbal2604/vehicle-tracking-api/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repo      *repositories.UserRepository
	tokenRepo *repositories.TokenBlacklistRepository
}

func NewUserService(repo *repositories.UserRepository, tokenRepo *repositories.TokenBlacklistRepository) UserService {
	return &UserServiceImpl{repo: repo, tokenRepo: tokenRepo}
}

func (s *UserServiceImpl) Register(name, email, password, role string) (*models.User, error) {
	if role == "" {
		role = "driver"
	}

	existingUser, _ := s.repo.FindByEmail(email)
	if existingUser != nil {
		return nil, errors.New("Email Already Registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}

	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserServiceImpl) Login(email, password string) (string, uint, error) {

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", 0, errors.New("Email Not Found")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", 0, errors.New("Wrong Password")
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		return "", 0, err
	}
	return token, user.ID, nil
}

func (s *UserServiceImpl) GetProfile(userID uint) (*models.User, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServiceImpl) UpdateProfile(userID uint, name, email string) (*models.User, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	user.Name = name
	user.Email = email

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServiceImpl) DeleteAccount(userID uint) error {
	return s.repo.Delete(userID)
}

func (s *UserServiceImpl) Logout(tokenString string, userID uint) error {
	// For simplicity, we set expiry to 24 hours. Ideally, it should be the same as the JWT expiry.
	expiresAt := time.Now().Add(24 * time.Hour)
	return s.tokenRepo.AddToken(tokenString, expiresAt)
}

func (s *UserServiceImpl) ListAllUsers() ([]models.User, error) {
	return s.repo.FindAllNonAdmin()
}
