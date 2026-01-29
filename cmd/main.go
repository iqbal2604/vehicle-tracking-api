package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/config"
	"github.com/iqbal2604/vehicle-tracking-api/handlers"
	"github.com/iqbal2604/vehicle-tracking-api/repositories"
	"github.com/iqbal2604/vehicle-tracking-api/routes"
	"github.com/iqbal2604/vehicle-tracking-api/services"
)

func main() {
	err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("DB Connection Failed")
	}

	app := fiber.New()

	userRepo := &repositories.UserRepository{}
	userService := services.NewUserService(userRepo)
	authHandler := handlers.NewAuthHandler(userService)

	routes.AuthRoutes(app, authHandler)

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Listen("localhost:3000")
}
