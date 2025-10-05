package main

import (
	"FreeLib/internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/health", handlers.HealthHandler).Methods("GET")
	r.HandleFunc("/api/books", handlers.GetBooksHandler).Methods("GET")
	r.HandleFunc("/api/book", handlers.GetByIDHandler).Methods("GET")
	r.HandleFunc("/api/search", handlers.SearchHandler).Methods("GET")
	r.HandleFunc("/api/create", handlers.CreateHandler).Methods("POST")
	r.HandleFunc("/api/books", handlers.DeleteHandler).Methods("DELETE")
	log.Println("FreeLib server sterting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}