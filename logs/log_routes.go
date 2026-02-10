package logs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/middlewares"
)

func LogRoute(router fiber.Router, logHandler *LogHandler) {
	protected := router.Group("/logs", middlewares.JWTMiddleware())

	protected.Get("/", logHandler.GetRecentLogs)
}
