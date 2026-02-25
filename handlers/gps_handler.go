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

	//Ambil query parameter start dan end
	start := c.Query("start")
	end := c.Query("end")

	// Check if user is admin
	user, err := h.userRepo.FindByID(userID)
	if err != nil {
		return helpers.ErrorResponse(c, 500, "Failed to check user role")
	}

	var locations []models.GPSLocation
	if user.Role == "admin" {
		locations, err = h.service.GetHistoryAdmin(vehicleID, start, end)
	} else {
		locations, err = h.service.GetHistory(userID, vehicleID, start, end)
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

func (h *GPSHandler) CreateGeofence(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var req dtos.GeofenceRequest

	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, 400, "Input tidak valid")
	}

	geofence := models.Geofence{
		Name:      req.Name,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Radius:    req.Radius,
		Type:      req.Type,
		UserID:    userID,
	}

	if err := h.service.CreateGeofence(&geofence); err != nil {
		return helpers.ErrorResponse(c, 500, err.Error())
	}

	return helpers.SuccessResponse(c, dtos.ToGeofenceResponse(geofence))
}

func (h *GPSHandler) ListGeofences(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	geofences, err := h.service.ListGeofences(userID)

	if err != nil {
		return helpers.ErrorResponse(c, 500, err.Error())
	}

	return helpers.SuccessResponse(c, dtos.ToGeofenceListResponse(geofences))
}

func (h *GPSHandler) DeleteGeofences(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, err := c.ParamsInt("id")

	if err != nil {
		return helpers.ErrorResponse(c, 400, "ID tidak Valid")
	}

	if err := h.service.DeleteGeofence(uint(id), userID); err != nil {
		return helpers.ErrorResponse(c, 500, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.Map{"message": "Geofence berhasil dihapus"})
}
