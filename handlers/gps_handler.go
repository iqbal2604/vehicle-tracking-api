package handlers

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/dtos"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
	models "github.com/iqbal2604/vehicle-tracking-api/models/domain"
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
	dto := dtos.ToGPSResponse(loc)
	return helpers.SuccessResponse(c, dto)
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
	locations, err := h.service.GetHistory(userID, vehicleID)
	if err != nil {
		return helpers.ErrorResponse(c, 404, err.Error())
	}

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

func (h *GPSHandler) StreamLocation(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	id, err := strconv.ParseUint(c.Params("vehicle_id"), 10, 64)
	if err != nil {
		return helpers.ErrorResponse(c, 400, "Invalid Vehicle ID")
	}

	vehicleID := uint(id)

	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")

	for {
		loc, err := h.service.GetLastLocation(userID, vehicleID)
		if err == nil {
			dto := dtos.ToGPSResponse(*loc)
			jsonData, _ := json.Marshal(dto)
			c.Write([]byte("data: " + string(jsonData) + "\n\n"))
		}
		time.Sleep(2 * time.Second)
	}
}
