package handlers

import (
	"io"
	"net/http"

	"example.auto.wiring/loggerfx"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(NewHandler))

func NewHandler(log loggerfx.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handler called")
		io.WriteString(w, "Hello, World!\n")
	})
}
