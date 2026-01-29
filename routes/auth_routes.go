package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/handlers"
)

func AuthRoutes(app *fiber.App, authHandler *handlers.AuthHandler) {
	app.Post("/api/register", authHandler.Register)
}
