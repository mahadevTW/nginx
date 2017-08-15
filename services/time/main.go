package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"time"
)

func NowHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().String()
	w.Write([]byte(currentTime))
}

func YesterdayHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Monday.String()
	w.Write([]byte(currentTime))
	w.Write([]byte(" Go to Work..!!!"))
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to time service , Use /now , /monday"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/now", NowHandler)
	r.HandleFunc("/monday", YesterdayHandler)
	r.HandleFunc("/", TimeHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8002", r))
}