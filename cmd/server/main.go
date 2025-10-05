package main

import (
	"FreeLib/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/health", handlers.HealthHandler)
	http.HandleFunc("/api/books", handlers.GetBooksHandler)
	http.HandleFunc("/api/book", handlers.GetByIDHandler)
	http.HandleFunc("/api/search", handlers.SearchHendler)
	http.HandleFunc("/api/create", handlers.CreateHeandler)
	log.Println("FreeLib server sterting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}