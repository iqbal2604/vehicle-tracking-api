package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/dtos"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
	"github.com/iqbal2604/vehicle-tracking-api/logs"
	models "github.com/iqbal2604/vehicle-tracking-api/models/domain"
	"github.com/iqbal2604/vehicle-tracking-api/repositories"
	"github.com/iqbal2604/vehicle-tracking-api/services"
)

type VehicleHandler struct {
	service    services.VehicleService
	repo       *repositories.UserRepository
	logService logs.LogService
}

func NewVehicleHandler(service services.VehicleService, repo *repositories.UserRepository, logService logs.LogService) *VehicleHandler {
	return &VehicleHandler{service: service, repo: repo, logService: logService}
}

func (h *VehicleHandler) getUserRole(userID uint) (string, error) {
	user, err := h.repo.FindByID(userID)
	if err != nil {
		return "", err
	}
	return user.Role, nil
}

func (h *VehicleHandler) CreateVehicle(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	role, err := h.getUserRole(userID)
	if err != nil {
		return helpers.ErrorResponse(c, 500, "Failed to get user role")
	}

	var v models.Vehicle
	if err := c.BodyParser(&v); err != nil {
		return helpers.Error(c, err.Error())
	}

	if role == "driver" {
		v.UserID = userID
	}

	if err := h.service.CreateVehicle(userID, &v); err != nil {
		return helpers.Error(c, err.Error())
	}

	ip := c.IP()
	h.logService.LogAuth("create_vehicle", &userID, "Vehicle Created:"+v.Name, ip)

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

	role, err := h.getUserRole(userID)
	if err != nil {
		return helpers.ErrorResponse(c, 500, "Failed to get user role")
	}

	v, err := h.service.GetVehicleByID(userID, uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, 404, "Vehicle Not Found")
	}

	if role == "driver" && v.UserID != userID {
		return helpers.ErrorResponse(c, 403, "Access Denied")
	}
	dto := dtos.ToVehicleResponse(*v)
	return helpers.SuccessResponse(c, dto)
}

func (h *VehicleHandler) ListVehicle(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	role, err := h.getUserRole(userID)
	if err != nil {
		return helpers.ErrorResponse(c, 500, "Failed to get user role")
	}

	var vehicles []models.Vehicle
	if role == "admin" {
		vehicles, err = h.service.ListAllVehicles()
	} else {
		vehicles, err = h.service.ListVehiclesByUser(userID)
	}

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

	role, err := h.getUserRole(userID)
	if err != nil {
		return helpers.ErrorResponse(c, 500, "Failed to get user role")
	}

	var v models.Vehicle
	if err := c.BodyParser(&v); err != nil {
		return helpers.ErrorResponse(c, 400, "Invalid Payload")
	}

	v.ID = uint(id)

	// Check ownership before updating
	existingVehicle, err := h.service.GetVehicleByID(userID, uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, 404, "Vehicle Not Found")
	}
	if role == "driver" && existingVehicle.UserID != userID {
		return helpers.ErrorResponse(c, 403, "Access Denied")
	}

	if err := h.service.UpdateVehicle(userID, &v); err != nil {
		return helpers.ErrorResponse(c, 400, err.Error())
	}

	ip := c.IP()
	h.logService.LogAuth("update_vehicle", &userID, "Vehicle Updated: ID"+strconv.Itoa(int(v.ID)), ip)

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

	ip := c.IP()
	h.logService.LogAuth("delete_vehicle", &userID, "Vehicle Deleted: ID "+idParam, ip)

	return helpers.SuccessResponse(c, fiber.Map{
		"message": "Vehicle Deleted",
	})
}
