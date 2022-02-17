package main

import (
	"example.auto.wiring/handlers"
	"example.auto.wiring/infra/configs"
	"example.auto.wiring/infra/database"
	"example.auto.wiring/infra/loggerfx"
	"example.auto.wiring/infra/serverfx"
	"example.auto.wiring/service/person"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		serverfx.Module,
		loggerfx.Module,
		handlers.Module,
		person.Module,
		database.Module,
		configs.Module,
	).Run()
}
