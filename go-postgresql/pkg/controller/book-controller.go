package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jangirvipin/go-postgresql/pkg/model"
	"github.com/jangirvipin/go-postgresql/pkg/utils"
	"net/http"
	"strconv"
)

func GetAllBook(w http.ResponseWriter, r *http.Request) {
	books, err := model.GetAllBooks()
	if err != nil {
		utils.SendError(w, 404, err.Error())
		return
	}
	utils.SendResponse(w, 200, books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ID, _ := strconv.ParseInt(id, 0, 0)
	book, err := model.GetBookById(ID)
	if err != nil {
		utils.SendError(w, 404, err.Error())
		return
	}
	utils.SendResponse(w, 200, book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ID, _ := strconv.ParseInt(id, 0, 0)
	book, err := model.DeleteBook(ID)

	if err != nil {
		utils.SendError(w, 404, err.Error())
		return
	}
	utils.SendResponse(w, 200, book)
	return
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &model.Book{}
	err := json.NewDecoder(r.Body).Decode(CreateBook)
	if err != nil {
		utils.SendError(w, 400, "Invalid request payload")
		return
	}
	createBook, err2 := CreateBook.CreateBook()
	if err2 != nil {
		utils.SendError(w, 500, err2.Error())
		return
	}
	utils.SendResponse(w, 201, createBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var UpdateBook = &model.Book{}
	err := json.NewDecoder(r.Body).Decode(UpdateBook)
	if err != nil {
		utils.SendError(w, 400, "Invalid request payload")
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	ID, _ := strconv.ParseInt(id, 0, 0)
	existingBook, err := model.GetBookById(ID)
	if err != nil {
		utils.SendError(w, 404, "Invlaid book ID")
		return
	}

	if existingBook.ID != UpdateBook.ID {
		utils.SendError(w, 400, "Something went wrong")
		return
	}

	if UpdateBook.Name != "" {
		existingBook.Name = UpdateBook.Name
	}
	if UpdateBook.Number != "" {
		existingBook.Number = UpdateBook.Number
	}
	if UpdateBook.Publisher != "" {
		existingBook.Publisher = UpdateBook.Publisher
	}
	if UpdateBook.Author != (model.Author{}) {
		if existingBook.Author.ID != 0 {
			existingBook.Author.Name = UpdateBook.Author.Name
			existingBook.Author.Email = UpdateBook.Author.Email
		}
	}

	updateBook, err := model.UpdateBook(existingBook)
	if err != nil {
		utils.SendError(w, 500, err.Error())
	}
	utils.SendResponse(w, 200, updateBook)
}
