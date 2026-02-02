package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/handlers"
	"github.com/iqbal2604/vehicle-tracking-api/middlewares"
)

func AuthRoutes(app *fiber.App, authHandler *handlers.AuthHandler) {
	app.Post("/api/register", authHandler.Register)
	app.Post("/api/login", authHandler.Login)
}

func UserRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	// Apply JWT middleware directly to specific route
	app.Get("/api/profile", middlewares.JWTMiddleware(), userHandler.GetProfile)
}

func VehicleRoutes(app *fiber.App, vehicleHandler *handlers.VehicleHandler) {
	protected := app.Group("/api/vehicles", middlewares.JWTMiddleware())

	protected.Post("", vehicleHandler.CreateVehicle)
	protected.Get("", vehicleHandler.ListVehicle)
	protected.Get("/:id", vehicleHandler.GetVehicle)
	protected.Put("/:id", vehicleHandler.UpdateVehicle)
	protected.Delete("/:id", vehicleHandler.DeleteVehicle)
}
