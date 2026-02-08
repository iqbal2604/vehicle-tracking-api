package dtos

import "github.com/iqbal2604/vehicle-tracking-api/models"

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"Name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func ToUserResponse(u models.User) UserResponse {
	return UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Role:  u.Role,
	}
}
