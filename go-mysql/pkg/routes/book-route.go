package routes

import (
	"github.com/gorilla/mux"
	"github.com/jangirvipin/go-mysql/pkg/controller"
)

var RegisterBookRoutes = func(r *mux.Router) {
	r.HandleFunc("/book", controller.CreateBook).Methods("POST")
	r.HandleFunc("/book", controller.GetBooks).Methods("GET")
	r.HandleFunc("/book/{id}", controller.GetBookById).Methods("GET")
	r.HandleFunc("/book/{id}", controller.UpdateBookById).Methods("PUT")
	r.HandleFunc("/book/{id}", controller.DeleteBookById).Methods("DELETE")
}
