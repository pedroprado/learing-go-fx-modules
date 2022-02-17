package handlers

import (
	"fmt"
	"net/http"

	"example.auto.wiring/service"
	"github.com/gorilla/mux"
)

type HttpHandlers struct {
	PersonHandler PersonHandler
}

func newHttpHandlers(personService *service.PersonService) *HttpHandlers {

	handlerPersonCreate := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		err := personService.Create()
		if err != nil {
			panic(err)
		}
		return
	}

	handlerPersonGet := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		vars := mux.Vars(r)
		id := vars["id"]

		fmt.Println("searched id: ", id)
		if id == "" {
			http.Error(w, "Id cannot be empty", http.StatusMethodNotAllowed)
			return
		}
		err := personService.Get(id)
		if err != nil {
			panic(err)
		}
		return
	}
	return &HttpHandlers{
		PersonHandler: PersonHandler{Create: handlerPersonCreate, Get: handlerPersonGet},
	}
}

type PersonHandler struct {
	Create http.HandlerFunc
	Get    http.HandlerFunc
}
