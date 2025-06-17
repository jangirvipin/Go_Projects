package main

import (
	"github.com/gorilla/mux"
	"github.com/jangirvipin/go-mysql/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
