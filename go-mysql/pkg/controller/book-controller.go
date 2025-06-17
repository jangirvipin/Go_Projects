package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jangirvipin/go-mysql/pkg/model"
	"log"
	"net/http"
	"strconv"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := model.GetBooks()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "Parameter id is missing", http.StatusBadRequest)
		return
	}
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	books, _ := model.GetBookById(ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &model.Book{}
	err := json.NewDecoder(r.Body).Decode(CreateBook)
	if err != nil {
		return
	}
	createBook := CreateBook.CreateBook()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createBook)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ID, _ := strconv.ParseInt(id, 0, 0)

	book := model.DeleteBook(ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var UpdateBook = &model.Book{}
	json.NewDecoder(r.Body).Decode(UpdateBook)
	vars := mux.Vars(r)
	id := vars["id"]

	ID, _ := strconv.ParseInt(id, 0, 0)
	existingBook, _ := model.GetBookById(ID)

	if UpdateBook.Name != "" {
		existingBook.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		existingBook.Author = UpdateBook.Author
	}
	if UpdateBook.Publisher != "" {
		existingBook.Publisher = UpdateBook.Publisher
	}

	updatedBook := model.UpdateBook(existingBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBook)
}
