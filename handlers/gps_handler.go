package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/dtos"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
	"github.com/iqbal2604/vehicle-tracking-api/logs"
	"github.com/iqbal2604/vehicle-tracking-api/models"
	"github.com/iqbal2604/vehicle-tracking-api/repositories"
	"github.com/iqbal2604/vehicle-tracking-api/services"
)

type GPSHandler struct {
	service    services.GPSService
	userRepo   *repositories.UserRepository
	logService logs.LogService
}

func NewGPSHandler(service services.GPSService, userRepo *repositories.UserRepository, logService logs.LogService) *GPSHandler {
	return &GPSHandler{service: service, userRepo: userRepo, logService: logService}
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

	ip := c.IP()
	h.logService.LogAuth("create_location", &userID, "GPS location created for vehicle", ip)

	dto := dtos.ToGPSResponse(loc)
	return helpers.SuccessResponse(c, dto)
}

func (h *GPSHandler) GetLastLocation(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	id, err := strconv.ParseUint(c.Params("vehicle_id"), 10, 64)
	if err != nil {
		return helpers.ErrorResponse(c, 400, "Invalid Vehicle ID")
	}

	// Check if user is admin
	user, err := h.userRepo.FindByID(userID)
	if err != nil {
		return helpers.ErrorResponse(c, 500, "Failed to check user role")
	}

	var loc *models.GPSLocation
	if user.Role == "admin" {
		loc, err = h.service.GetLastLocationAdmin(uint(id))
	} else {
		loc, err = h.service.GetLastLocation(userID, uint(id))
	}

	if err != nil {
		return helpers.ErrorResponse(c, 400, err.Error())
	}
	dto := dtos.ToGPSResponse(*loc)
	return helpers.SuccessResponse(c, dto)
}

func (h *GPSHandler) GetHistory(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	id, err := strconv.ParseUint(c.Params("vehicle_id"), 10, 64)
	if err != nil {
		return helpers.ErrorResponse(c, 400, "Invalid Vehicle ID")
	}
	vehicleID := uint(id)

	// Check if user is admin
	user, err := h.userRepo.FindByID(userID)
	if err != nil {
		return helpers.ErrorResponse(c, 500, "Failed to check user role")
	}

	var locations []models.GPSLocation
	if user.Role == "admin" {
		locations, err = h.service.GetHistoryAdmin(vehicleID)
	} else {
		locations, err = h.service.GetHistory(userID, vehicleID)
	}

	if err != nil {
		return helpers.ErrorResponse(c, 404, err.Error())
	}

	ip := c.IP()
	h.logService.LogAuth("get_history_location", &userID, "Get location history for vehicle", ip)

	var result []dtos.GPSResponse
	for _, g := range locations {
		result = append(result, dtos.ToGPSResponse(g))
	}

	response := dtos.GPSHistoryResponse{
		VehicleID: vehicleID,
		Locations: result,
	}
	return helpers.SuccessResponse(c, response)
}
