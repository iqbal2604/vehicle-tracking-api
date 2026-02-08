package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/iqbal2604/vehicle-tracking-api/config"
	"github.com/iqbal2604/vehicle-tracking-api/routes"
)

func main() {
	err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("DB Connection Failed")
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Initialize handlers using Wire
	authHandler := InitializeAuthHandler()
	userHandler := InitializeUserHandler()
	vehicleHandler := InitializeVehicleHandler()
	gpsHandler, hub := InitializedGPSHandler()

	go hub.Run()

	//Group
	api := app.Group("/api")

	//Routes
	routes.VehicleRoutes(api, vehicleHandler)
	routes.AuthRoutes(api, authHandler)
	routes.UserRoutes(api, userHandler)
	routes.GPSRoute(api, gpsHandler)
	routes.WebsocketRoutes(app, hub)

	app.Listen("localhost:3000")
}
