package main

import (
	"example.manual.wiring/handlers"
	"example.manual.wiring/logger"
	"example.manual.wiring/server"
)

func main() {
	logger := logger.NewLogger()
	handler := handlers.NewHandler(logger)
	server.RegisterHandlers(handler)
	server.StartServer()
}
