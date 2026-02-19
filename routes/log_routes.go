package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
	"github.com/iqbal2604/vehicle-tracking-api/logs"
	"github.com/redis/go-redis/v9"
)

func LogRoute(router fiber.Router, logHandler *logs.LogHandler, rdb *redis.Client) {
	protected := router.Group("/logs", helpers.JWTMiddleware(rdb))

	protected.Get("/", logHandler.GetRecentLogs)
}
