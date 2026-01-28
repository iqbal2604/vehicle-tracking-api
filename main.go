package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbal2604/vehicle-tracking-api/routes"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
