package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/handlers"
)

func AuthRoutes(router fiber.Router, authHandler *handlers.AuthHandler) {
	router.Post("/register", authHandler.Register)
	router.Post("/login", authHandler.Login)
}
