package main

import (
	"FreeLib/internal/handlers"
	"FreeLib/pkg/config"
	"FreeLib/pkg/database"
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()
	 ctx := context.Background()
	log.Println(cfg)
	pool, err := database.ConnectDB(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	log.Println("Connect db")
	
	r := mux.NewRouter()
	r.HandleFunc("/api/health", handlers.HealthHandler).Methods("GET")
	r.HandleFunc("/api/books", handlers.GetBooksHandler).Methods("GET")
	r.HandleFunc("/api/book", handlers.GetByIDHandler).Methods("GET")
	r.HandleFunc("/api/search", handlers.SearchHandler).Methods("GET")
	r.HandleFunc("/api/create", handlers.CreateHandler).Methods("POST")
	r.HandleFunc("/api/books", handlers.DeleteHandler).Methods("DELETE")

	r.HandleFunc("/api/register", handlers.RegisterHandler).Methods("POST")

	log.Println("FreeLib server sterting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}