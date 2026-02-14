package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/config"
	"github.com/iqbal2604/vehicle-tracking-api/handlers"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
)

func VehicleRoutes(router fiber.Router, vehicleHandler *handlers.VehicleHandler) {
	protected := router.Group("/vehicles", helpers.JWTMiddleware(config.DB))

	protected.Post("", vehicleHandler.CreateVehicle)
	protected.Get("", vehicleHandler.ListVehicle)
	protected.Get("/user/:userId", vehicleHandler.ListVehiclesByUserID)
	protected.Get("/:id", vehicleHandler.GetVehicle)
	protected.Put("/:id", vehicleHandler.UpdateVehicle)
	protected.Delete("/:id", vehicleHandler.DeleteVehicle)
}
