//go:build wireinject
// +build wireinject

// gobuild wireinject
package main

import (
	"github.com/google/wire"
	"github.com/iqbal2604/vehicle-tracking-api/config"
	"github.com/iqbal2604/vehicle-tracking-api/handlers"
	"github.com/iqbal2604/vehicle-tracking-api/logs"
	"github.com/iqbal2604/vehicle-tracking-api/repositories"
	"github.com/iqbal2604/vehicle-tracking-api/services"
	"github.com/iqbal2604/vehicle-tracking-api/websocket"
)

type GPSComponents struct {
	Handler *handlers.GPSHandler
	Hub     *websocket.Hub
}

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
		logs.NewLogRepository,
		logs.NewLogServiceImpl,
		repositories.NewUserRepository,
		services.NewUserService,
		handlers.NewAuthHandler,
	)
	return nil
}

func InitializeVehicleHandler() *handlers.VehicleHandler {
	wire.Build(
		config.NewDatabase,
		logs.NewLogRepository,
		logs.NewLogServiceImpl,
		repositories.NewUserRepository,
		repositories.NewVehicleRepository,
		services.NewVehicleService,
		handlers.NewVehicleHandler,
	)
	return nil
}

func InitializedGPSHandler() *GPSComponents {
	wire.Build(

		config.NewDatabase,
		logs.NewLogRepository,
		logs.NewLogServiceImpl,
		repositories.NewGPSRepository,
		repositories.NewVehicleRepository,
		services.NewGPSService,
		handlers.NewGPSHandler,
		websocket.NewHub,
		wire.Struct(new(GPSComponents), "*"),
	)
	return nil

}
func InitializeLogHandler() *logs.LogHandler {
	wire.Build(
		config.NewDatabase,
		logs.NewLogRepository,
		logs.NewLogServiceImpl,
		logs.NewLogHandler,
	)
	return nil
}
