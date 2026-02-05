package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/handlers"
	"github.com/iqbal2604/vehicle-tracking-api/middlewares"
)

func VehicleRoutes(router fiber.Router, vehicleHandler *handlers.VehicleHandler) {
	protected := router.Group("/vehicles", middlewares.JWTMiddleware())

	protected.Post("", vehicleHandler.CreateVehicle)
	protected.Get("", vehicleHandler.ListVehicle)
	protected.Get("/:id", vehicleHandler.GetVehicle)
	protected.Put("/:id", vehicleHandler.UpdateVehicle)
	protected.Delete("/:id", vehicleHandler.DeleteVehicle)
}
