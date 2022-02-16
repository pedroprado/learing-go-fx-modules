package handlers

import (
	"example.manual.wiring/logger"
	"io"
	"net/http"
)

func NewHandler(log logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handler called")
		io.WriteString(w, "Hello, World!\n")
	})
}
