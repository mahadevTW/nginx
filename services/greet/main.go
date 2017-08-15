package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to greeting service..!\n USE /morning and /evening handlers"))
}

func morningHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Good morning..!\n"))
}

func eveningHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Good evening.. go to home :P !\n"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", greetHandler)
	r.HandleFunc("/morning", morningHandler)
	r.HandleFunc("/evening", eveningHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}