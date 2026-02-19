package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/handlers"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
	"github.com/redis/go-redis/v9"
)

func AuthRoutes(router fiber.Router, authHandler *handlers.AuthHandler, rdb *redis.Client) {
	router.Post("/register", authHandler.Register)
	router.Post("/login", authHandler.Login)
	router.Post("/logout", helpers.JWTMiddleware(rdb), authHandler.Logout)
}
