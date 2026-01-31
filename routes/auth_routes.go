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

func ProtectedRoutes(app *fiber.App) {
	protected := app.Group("/api/protected", middlewares.JWTMiddleware())

	protected.Get("/test", func(c *fiber.Ctx) error {
		userID := c.Locals("user_id")
		return c.JSON(fiber.Map{
			"user_id": userID,
		})
	})
}

func UserRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	// Apply JWT middleware directly to specific route
	app.Get("/api/profile", middlewares.JWTMiddleware(), userHandler.GetProfile)
}
