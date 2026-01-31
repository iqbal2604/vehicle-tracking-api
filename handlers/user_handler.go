package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
	"github.com/iqbal2604/vehicle-tracking-api/services"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Profile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	user, err := h.service.GetProfile(userID)
	if err != nil {
		return helpers.Error(c, err.Error())
	}

	return helpers.SuccessResponse(c, user)
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	user, err := h.service.GetProfile(userID)
	if err != nil {
		return helpers.ErrorResponse(c, 400, "User not found")
	}

	return helpers.SuccessResponse(c, fiber.Map{
		"message": "Profile retrieved",
		"data":    user,
	})
}
