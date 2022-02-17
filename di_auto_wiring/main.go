package main

import (
	// "example.auto.wiring/handlers"
	"example.auto.wiring/handlers"
	"example.auto.wiring/infra/configs"
	"example.auto.wiring/infra/database"
	"example.auto.wiring/infra/loggerfx"
	"example.auto.wiring/infra/repository"
	"example.auto.wiring/infra/serverfx"
	"example.auto.wiring/service"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		loggerfx.Module,
		configs.Module,
		database.Module,
		repository.Module,
		service.Module,
		handlers.Module,
		serverfx.Module,
	).Run()
}
