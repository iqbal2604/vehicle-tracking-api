package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/handlers"
	"github.com/iqbal2604/vehicle-tracking-api/middlewares"
)

func UserRoutes(router fiber.Router, userHandler *handlers.UserHandler) {
	// Apply JWT middleware directly to specific route
	router.Get("/profile", middlewares.JWTMiddleware(), userHandler.GetProfile)
}
