package handlers

import (
	"FreeLib/internal/repository"
	"encoding/json"
	"log"
	"net/http"
)

var myRepo = &repository.MockBookRepo{}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string {
		"status": "ok",
		"service": "FreeLib API",
		"version": "1.0.0",
	})

}

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := myRepo.GetAll()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}