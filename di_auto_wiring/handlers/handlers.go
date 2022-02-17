package handlers

import (
	"io"
	"net/http"

	"example.auto.wiring/infra/loggerfx"
	"example.auto.wiring/service/person"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(NewHandler))

func NewHandler(log loggerfx.Logger, personService *person.PersonService) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		personService.Create()
		log.Println("Handler called")
		io.WriteString(w, "Hello, World!\n")
	})
}
