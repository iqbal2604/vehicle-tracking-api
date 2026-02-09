package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
	"github.com/iqbal2604/vehicle-tracking-api/requests"
	"github.com/iqbal2604/vehicle-tracking-api/services"
)

type AuthHandler struct {
	service services.UserService
}

func NewAuthHandler(service services.UserService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {

	var req requests.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, 400, "Invalid Request")
	}

	err := h.service.Register(req.Name, req.Email, req.Password, req.Role)
	if err != nil {
		return helpers.ErrorResponse(c, 400, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.Map{
		"status":  201,
		"message": "Register Succses",
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req requests.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, 400, "Invalid Request")
	}

	token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return helpers.ErrorResponse(c, 400, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.Map{
		"message": "Login Success",
		"token":   token,
	})

}
