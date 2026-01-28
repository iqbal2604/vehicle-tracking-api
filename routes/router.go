package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.Ping)

}
