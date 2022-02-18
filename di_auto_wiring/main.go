package main

import (
	"example.auto.wiring/src/core/service"
	"example.auto.wiring/src/handlers"
	"example.auto.wiring/src/infra/configs"
	"example.auto.wiring/src/infra/database"
	"example.auto.wiring/src/infra/loggerfx"
	"example.auto.wiring/src/infra/repository"
	"example.auto.wiring/src/infra/serverfx"
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
