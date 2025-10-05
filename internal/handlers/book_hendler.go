package handlers

import (
	"FreeLib/internal/models"
	"FreeLib/internal/repository"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

func GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	idUint, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	log.Println(uint(idUint))
	book, err := myRepo.GetByID(uint(idUint))
	
	if err != nil {
		log.Println(err)
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func SearchHendler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	books, err := myRepo.Search(query)
	if err != nil {
		log.Println(err)
		http.Error(w, "Books not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func CreateHeandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := myRepo.Create(&book); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}