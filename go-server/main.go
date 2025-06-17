package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse Form error: %v", err)
	}

	fmt.Fprintf(w, "POST request received!\n")

	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")
	fmt.Fprintf(w, "Hello %s, %s", name, email)
	fmt.Fprintf(w, "Your message: %s", message)
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/func" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello from the server!")
}

func main() {

	fileServe := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServe)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/func", handleFunc)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Server running on port 8080\n")
	}
}
