package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/config"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
	"github.com/iqbal2604/vehicle-tracking-api/logs"
)

func LogRoute(router fiber.Router, logHandler *logs.LogHandler) {
	protected := router.Group("/logs", helpers.JWTMiddleware(config.DB))

	protected.Get("/", logHandler.GetRecentLogs)
}
