package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
	"github.com/iqbal2604/vehicle-tracking-api/logs"
	"github.com/iqbal2604/vehicle-tracking-api/requests"
	"github.com/iqbal2604/vehicle-tracking-api/services"
)

type AuthHandler struct {
	service    services.UserService
	logService logs.LogService
}

func NewAuthHandler(service services.UserService, logService logs.LogService) *AuthHandler {
	return &AuthHandler{service: service, logService: logService}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {

	var req requests.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, 400, "Invalid Request")
	}

	user, err := h.service.Register(req.Name, req.Email, req.Password, req.Role)
	if err != nil {
		return helpers.ErrorResponse(c, 400, err.Error())
	}

	ip := c.IP()
	h.logService.LogAuth("register", &user.ID, "User Registered:"+req.Email, ip)

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

	token, userID, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return helpers.ErrorResponse(c, 400, err.Error())
	}

	ip := c.IP()
	h.logService.LogAuth("login", &userID, "User Logged in: "+req.Email, ip)

	return helpers.SuccessResponse(c, fiber.Map{
		"message": "Login Success",
		"token":   token,
	})

}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	userID := c.Locals("user_id").(uint)

	err := h.service.Logout(tokenString, userID)
	if err != nil {
		return helpers.ErrorResponse(c, 500, "Logout Failed")
	}

	ip := c.IP()
	h.logService.LogAuth("logout", &userID, "User Logged Out", ip)

	return helpers.SuccessResponse(c, fiber.Map{
		"message": "Logout Success",
	})
}
