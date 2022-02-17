package main

import (
	"example.auto.wiring/handlers"
	"example.auto.wiring/loggerfx"
	"example.auto.wiring/serverfx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		serverfx.Module,
		loggerfx.Module,
		handlers.Module,
	).Run()
}
