package main

import (
	"FreeLib/internal/handlers"
	"FreeLib/internal/repository/postgres"
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
	pool, err := database.ConnectDB(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	log.Println("Connect db")
	
	bookRepo := postgres.NewBookRepository(pool)

	bookHandler := handlers.NewBookhandler(bookRepo)

	r := mux.NewRouter()
	r.HandleFunc("/api/health", bookHandler.HealthHandler).Methods("GET")
	r.HandleFunc("/api/books", bookHandler.GetBooksHandler).Methods("GET")
	r.HandleFunc("/api/book", bookHandler.GetByIDHandler).Methods("GET")
	r.HandleFunc("/api/search", bookHandler.SearchHandler).Methods("GET")
	r.HandleFunc("/api/create", bookHandler.CreateHandler).Methods("POST")
	r.HandleFunc("/api/books", bookHandler.DeleteHandler).Methods("DELETE")

	r.HandleFunc("/api/register", handlers.RegisterHandler).Methods("POST")

	log.Println("FreeLib server sterting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}