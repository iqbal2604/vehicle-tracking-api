package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/config"
	"github.com/iqbal2604/vehicle-tracking-api/routes"
)

func main() {
	config.ConnectDatabase()

	app := fiber.New()

	routes.SetupRoutes(app)

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
