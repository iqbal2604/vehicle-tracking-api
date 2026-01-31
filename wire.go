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
