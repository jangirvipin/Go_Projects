package routes

import (
	"github.com/gorilla/mux"
	"github.com/jangirvipin/go-postgresql/pkg/controller"
)

func RegisterBooksRoutes(router *mux.Router) {
	router.HandleFunc("/book", controller.GetAllBook).Methods("GET")
	router.HandleFunc("/book", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controller.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{id}", controller.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{id}", controller.UpdateBook).Methods("PUT")
}
