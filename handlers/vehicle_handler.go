package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/dtos"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
	"github.com/iqbal2604/vehicle-tracking-api/models"
	"github.com/iqbal2604/vehicle-tracking-api/services"
)

type VehicleHandler struct {
	service *services.VehicleService
}

func NewVehicleHandler(service *services.VehicleService) *VehicleHandler {
	return &VehicleHandler{service: service}
}

func (h *VehicleHandler) CreateVehicle(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var v models.Vehicle
	if err := c.BodyParser(&v); err != nil {
		return helpers.Error(c, err.Error())
	}

	if err := h.service.CreateVehicle(userID, &v); err != nil {
		return helpers.Error(c, err.Error())
	}

	dto := dtos.ToVehicleResponse(v)

	return helpers.SuccessResponse(c, fiber.Map{
		"message": "Vehicle Created",
		"data":    dto,
	})
}

func (h *VehicleHandler) GetVehicle(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return helpers.ErrorResponse(c, 400, "Inalid ID")
	}

	v, err := h.service.GetVehicleByID(userID, uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, 404, "Vehicle Not Found")
	}
	dto := dtos.ToVehicleResponse(*v)
	return helpers.SuccessResponse(c, dto)
}

func (h *VehicleHandler) ListVehicle(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	vehicles, err := h.service.ListVehiclesByUser(userID)
	if err != nil {
		return helpers.Error(c, err.Error())
	}
	var result []dtos.VehicleResponse

	for _, v := range vehicles {
		result = append(result, dtos.ToVehicleResponse(v))
	}
	return helpers.SuccessResponse(c, result)
}

func (h *VehicleHandler) UpdateVehicle(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return helpers.ErrorResponse(c, 400, "Invalid ID")
	}
	userID := c.Locals("user_id").(uint)

	var v models.Vehicle
	if err := c.BodyParser(&v); err != nil {
		return helpers.ErrorResponse(c, 400, "Invalid Payload")
	}

	v.ID = uint(id)

	if err := h.service.UpdateVehicle(userID, &v); err != nil {
		return helpers.ErrorResponse(c, 400, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.Map{
		"message": "Vehicle Updated",
		"data":    dtos.ToVehicleResponse(v),
	})

}

func (h *VehicleHandler) DeleteVehicle(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return helpers.ErrorResponse(c, 400, "Invalid ID")
	}

	if err := h.service.DeleteVehicle(userID, uint(id)); err != nil {
		return helpers.ErrorResponse(c, 400, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.Map{
		"message": "Vehicle Deleted",
	})
}
