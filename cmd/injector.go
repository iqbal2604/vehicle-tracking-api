//go:build wireinject
// +build wireinject

// gobuild wireinject

package main

import (
	"github.com/google/wire"
	"github.com/iqbal2604/vehicle-tracking-api/config"
	"github.com/iqbal2604/vehicle-tracking-api/handlers"
	"github.com/iqbal2604/vehicle-tracking-api/repositories"
	"github.com/iqbal2604/vehicle-tracking-api/services"
	"github.com/iqbal2604/vehicle-tracking-api/websocket"
)

func InitializeUserHandler() *handlers.UserHandler {
	wire.Build(
		config.NewDatabase,
		repositories.NewUserRepository,
		services.NewUserService,
		handlers.NewUserHandler,
	)
	return nil
}

func InitializeAuthHandler() *handlers.AuthHandler {
	wire.Build(
		config.NewDatabase,
		repositories.NewUserRepository,
		services.NewUserService,
		handlers.NewAuthHandler,
	)
	return nil
}

func InitializeVehicleHandler() *handlers.VehicleHandler {
	wire.Build(
		config.NewDatabase,
		repositories.NewUserRepository,
		repositories.NewVehicleRepository,
		services.NewVehicleService,
		handlers.NewVehicleHandler,
	)
	return nil
}

func InitializedGPSHandler() (*handlers.GPSHandler, *websocket.Hub) {
	wire.Build(

		config.NewDatabase,
		repositories.NewGPSRepository,
		repositories.NewVehicleRepository,
		services.NewGPSService,
		handlers.NewGPSHandler,
		websocket.NewHub,
	)
	return nil, nil

}
