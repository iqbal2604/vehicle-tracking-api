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
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	authHandler := handlers.NewAuthHandler(userService)
	userHandler := handlers.NewUserHandler(userService)

	vehicleRepo := repositories.NewVehicleRepository(db)
	vehicleService := services.NewVehicleService(vehicleRepo)
	vehicleHandler := handlers.NewVehicleHandler(vehicleService)

	routes.VehicleRoutes(app, vehicleHandler)
	routes.AuthRoutes(app, authHandler)
	routes.UserRoutes(app, userHandler)

	app.Listen("localhost:3000")
}
