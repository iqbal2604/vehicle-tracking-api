package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/config"
	"github.com/iqbal2604/vehicle-tracking-api/handlers"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
)

func UserRoutes(router fiber.Router, userHandler *handlers.UserHandler) {
	// Apply JWT middleware directly to specific route
	router.Get("/profile", helpers.JWTMiddleware(config.DB), userHandler.GetProfile)
}
