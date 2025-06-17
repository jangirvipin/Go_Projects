package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jangirvipin/go-postgresql/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBooksRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server")
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		fmt.Printf("Route: %s, Methods: %v\n", path, methods)
		return nil
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
