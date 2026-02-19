package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/handlers"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
	"github.com/redis/go-redis/v9"
)

func VehicleRoutes(router fiber.Router, vehicleHandler *handlers.VehicleHandler, rdb *redis.Client) {
	protected := router.Group("/vehicles", helpers.JWTMiddleware(rdb))

	protected.Post("", vehicleHandler.CreateVehicle)
	protected.Get("", vehicleHandler.ListVehicle)
	protected.Get("/user/:userId", vehicleHandler.ListVehiclesByUserID)
	protected.Get("/:id", vehicleHandler.GetVehicle)
	protected.Put("/:id", vehicleHandler.UpdateVehicle)
	protected.Delete("/:id", vehicleHandler.DeleteVehicle)
}
