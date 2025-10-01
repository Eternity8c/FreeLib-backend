package main

import (
	"FreeLib/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/health", handlers.HealthHandler)
	http.HandleFunc("/api/books", handlers.GetBooksHandler)
	log.Println("FreeLib server sterting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}