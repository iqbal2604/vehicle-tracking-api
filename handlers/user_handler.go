package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/dtos"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
	"github.com/iqbal2604/vehicle-tracking-api/requests"
	"github.com/iqbal2604/vehicle-tracking-api/services"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	user, err := h.service.GetProfile(userID)
	if err != nil {
		return helpers.ErrorResponse(c, 400, "User not found")
	}

	dto := dtos.ToUserResponse(*user)
	return helpers.SuccessResponse(c, fiber.Map{
		"message": "Profile retrieved",
		"data":    dto,
	})
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.service.ListAllUsers()
	if err != nil {
		return helpers.ErrorResponse(c, 500, "Failed to list users")
	}

	var result []dtos.UserResponse
	for _, u := range users {
		result = append(result, dtos.ToUserResponse(u))
	}

	return helpers.SuccessResponse(c, fiber.Map{
		"message": "Users retrieved",
		"data":    result,
	})
}
func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var req requests.RegisterRequest // Reusing RegisterRequest for name/email
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, 400, "Invalid payload")
	}

	user, err := h.service.UpdateProfile(userID, req.Name, req.Email)
	if err != nil {
		return helpers.ErrorResponse(c, 400, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.Map{
		"message": "Profile updated",
		"data":    dtos.ToUserResponse(*user),
	})
}

func (h *UserHandler) DeleteAccount(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	if err := h.service.DeleteAccount(userID); err != nil {
		return helpers.ErrorResponse(c, 400, "Failed to delete account")
	}

	return helpers.SuccessResponse(c, fiber.Map{
		"message": "Account deleted successfully",
	})
}
