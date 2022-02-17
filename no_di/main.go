package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

var logger = log.New(os.Stdout, "[Me]", 0)

func main() {

	http.Handle("/", http.HandlerFunc(helloWorldHandler))

	http.ListenAndServe(":8090", nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	logger.Println("Handler called")
	io.WriteString(w, "Hello, World!\n")

}
