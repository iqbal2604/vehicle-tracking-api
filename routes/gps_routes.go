package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/config"
	"github.com/iqbal2604/vehicle-tracking-api/handlers"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
)

func GPSRoute(router fiber.Router, gpsHandler *handlers.GPSHandler) {
	protected := router.Group("/gps", helpers.JWTMiddleware(config.DB))

	protected.Post("/", gpsHandler.CreateLocation)
	protected.Get("/history/:vehicle_id", gpsHandler.GetHistory)
	protected.Get("/last/:vehicle_id", gpsHandler.GetLastLocation)
}
