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

	db := config.NewDatabase()

	//Repository
	userRepo := repositories.NewUserRepository(db)
	vehicleRepo := repositories.NewVehicleRepository(db)
	gpsRepo := repositories.NewGPSRepository(db)
	//Services
	userService := services.NewUserService(userRepo)
	vehicleService := services.NewVehicleService(vehicleRepo)
	gpsService := services.NewGPSService(gpsRepo, vehicleRepo)
	//Handlers
	authHandler := handlers.NewAuthHandler(userService)
	userHandler := handlers.NewUserHandler(userService)
	vehicleHandler := handlers.NewVehicleHandler(vehicleService)
	gpsHandler := handlers.NewGPSHandler(gpsService)

	//Group
	api := app.Group("/api")

	//Routes
	routes.VehicleRoutes(api, vehicleHandler)
	routes.AuthRoutes(api, authHandler)
	routes.UserRoutes(api, userHandler)
	routes.GPSRoute(api, gpsHandler)

	app.Listen("localhost:3000")
}
