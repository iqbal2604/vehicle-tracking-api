package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
	"github.com/iqbal2604/vehicle-tracking-api/models"
	"github.com/iqbal2604/vehicle-tracking-api/services"
)

type GPSHandler struct {
	service *services.GPSService
}

func NewGPSHandler(service *services.GPSService) *GPSHandler {
	return &GPSHandler{service: service}
}

func (h *GPSHandler) CreateLocation(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var loc models.GPSLocation
	if err := c.BodyParser(&loc); err != nil {
		return helpers.ErrorResponse(c, 400, "Invalid Payload")
	}

	if err := h.service.CreateLocation(userID, &loc); err != nil {
		return helpers.ErrorResponse(c, 400, err.Error())
	}
	return helpers.SuccessResponse(c, loc)
}

func (h *GPSHandler) GetLastLocation(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	id, err := strconv.ParseUint(c.Params("vehicle_id"), 10, 64)
	if err != nil {
		return helpers.ErrorResponse(c, 400, "Invalid Vehicle ID")
	}

	loc, err := h.service.GetLastLocation(userID, uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, 400, err.Error())
	}
	return helpers.SuccessResponse(c, loc)
}

func (h *GPSHandler) GetHistory(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	id, err := strconv.ParseUint(c.Params("vehicle_id"), 10, 64)
	if err != nil {
		return helpers.ErrorResponse(c, 400, "Invalid Vehicle ID")
	}

	locations, err := h.service.GetHistory(userID, uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, 404, err.Error())
	}
	return helpers.SuccessResponse(c, locations)
}
