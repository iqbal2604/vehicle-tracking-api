package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/handlers"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
	"github.com/redis/go-redis/v9"
)

func GPSRoute(router fiber.Router, gpsHandler *handlers.GPSHandler, rdb *redis.Client) {
	protected := router.Group("/gps", helpers.JWTMiddleware(rdb))

	protected.Post("/", gpsHandler.CreateLocation)
	protected.Get("/history/:vehicle_id", gpsHandler.GetHistory)
	protected.Get("/last/:vehicle_id", gpsHandler.GetLastLocation)
}
