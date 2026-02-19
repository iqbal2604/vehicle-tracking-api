package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/handlers"
	"github.com/iqbal2604/vehicle-tracking-api/helpers"
	"github.com/redis/go-redis/v9"
)

func UserRoutes(router fiber.Router, userHandler *handlers.UserHandler, rdb *redis.Client) {
	protected := router.Group("/profile", helpers.JWTMiddleware(rdb))
	protected.Get("", userHandler.GetProfile)
	protected.Put("", userHandler.UpdateProfile)
	protected.Delete("", userHandler.DeleteAccount)

	router.Get("/users", helpers.JWTMiddleware(rdb), userHandler.ListUsers)
}
