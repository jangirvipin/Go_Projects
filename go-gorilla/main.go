package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movie

func handleMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		log.Println("Something went wrong", err)
		return
	}
}

func handleMoviesByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	params := mux.Vars(r)
	id := params["id"]

	for _, movie := range movies {
		if movie.ID == id {
			err := json.NewEncoder(w).Encode(movie)
			if err != nil {
				log.Println("Something went wrong", err)
				return
			}
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func handleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	movies = append(movies, movie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(movies)

	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}

func handleMovieDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	id := params["id"]

	for index, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}

	http.Error(w, "Movie not found", http.StatusNotFound)
}

func handleMovieUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	var updatedMovie Movie

	err := json.NewDecoder(r.Body).Decode(&updatedMovie)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	for index, movie := range movies {
		if movie.ID == id {
			movies[index] = updatedMovie
			movies[index].ID = id
			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(movies)
			if err != nil {
				http.Error(w, "Something went wrong", http.StatusInternalServerError)
			}
		}
	}

	http.Error(w, "Movie not found", http.StatusNotFound)

}

func main() {

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "123456",
		Title: "Movie One",
		Director: &Director{
			FirstName: "John",
			LastName:  "Doe",
		}})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "654321",
		Title: "Movie Two",
		Director: &Director{
			FirstName: "Jane",
			LastName:  "Doe",
		}})

	r := mux.NewRouter()

	r.HandleFunc("/movies", handleMovies).Methods("GET")
	r.HandleFunc("/movies", handleCreate).Methods("POST")
	r.HandleFunc("/movies/{id}", handleMoviesByID).Methods("GET")
	r.HandleFunc("/movies/{id}", handleMovieUpdate).Methods("PUT")
	r.HandleFunc("/movies/{id}", handleMovieDelete).Methods("DELETE")

	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
