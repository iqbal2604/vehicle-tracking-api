package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/iqbal2604/vehicle-tracking-api/config"
	"github.com/iqbal2604/vehicle-tracking-api/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, continuing with environment variables")
	}

	if err := config.ConnectDatabase(); err != nil {
		log.Fatal("DB Connection Failed")
	}

	rdb := config.NewRedisClient()
	defer rdb.Close()

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
	gpsApp := InitializedGPSHandler()
	logHandler := InitializeLogHandler()

	go gpsApp.Hub.Run()

	//Group
	api := app.Group("/api")

	//Routes
	routes.VehicleRoutes(api, vehicleHandler, rdb)
	routes.AuthRoutes(api, authHandler, rdb)
	routes.UserRoutes(api, userHandler, rdb)
	routes.GPSRoute(api, gpsApp.Handler, rdb)
	routes.LogRoute(api, logHandler, rdb)
	routes.WebsocketRoutes(app, gpsApp.Hub)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
