package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"example.auto.wiring/core/domain"
	"example.auto.wiring/core/service"
	"github.com/gorilla/mux"
)

type HttpHandlers struct {
	PersonHandler PersonHandler
}

func newHttpHandlers(personService *service.PersonService) *HttpHandlers {

	handlerPersonCreate := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		personRequest := domain.Person{}
		if err := json.Unmarshal(body, &personRequest); err != nil {
			http.Error(w, "not a person", http.StatusBadRequest)
			return
		}

		created, err := personService.Create(personRequest)
		if err != nil {
			panic(err)
		}
		response, err := json.Marshal(created)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return
	}

	handlerPersonGet := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		vars := mux.Vars(r)
		id := vars["id"]

		if id == "" {
			http.Error(w, "Id cannot be empty", http.StatusMethodNotAllowed)
			return
		}

		person, err := personService.Get(id)
		if err != nil {
			panic(err)
		}
		if person == nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		response, err := json.Marshal(person)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
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
