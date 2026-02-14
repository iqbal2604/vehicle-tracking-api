package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/config"
	"github.com/iqbal2604/vehicle-tracking-api/handlers"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
)

func AuthRoutes(router fiber.Router, authHandler *handlers.AuthHandler) {
	router.Post("/register", authHandler.Register)
	router.Post("/login", authHandler.Login)
	router.Post("/logout", helpers.JWTMiddleware(config.DB), authHandler.Logout)
}
