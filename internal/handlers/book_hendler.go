package handlers

import (
	"FreeLib/internal/models"
	"FreeLib/internal/repository"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type BookHandler struct {
	repo repository.BookRepository
}

func NewBookHandler(repo repository.BookRepository) *BookHandler {
	return &BookHandler{
		repo: repo,
	}
}

func (h *BookHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string {
		"status": "ok",
		"service": "FreeLib API",
		"version": "1.0.0",
	})

}

func (h *BookHandler) GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := h.repo.GetAll()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	idUint, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	log.Println(uint(idUint))
	book, err := h.repo.GetByID(uint(idUint))
	
	if err != nil {
		log.Println(err)
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	books, err := h.repo.Search(query)
	if err != nil {
		log.Println(err)
		http.Error(w, "Books not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&book); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	strID := r.URL.Query().Get("id")
	intId, err := strconv.Atoi(strID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(uint(intId)); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}