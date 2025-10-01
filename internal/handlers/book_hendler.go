package handlers

import (
	"encoding/json"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string {
		"status": "ok",
		"service": "FreeLib API",
		"version": "1.0.0",
	})
}

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books := []map[string]interface{} {
		{
			"id": 1,
			"title": "Мастер и Маргарита",
			"author": "М.А. Булгаков",
			"genre": "Классика",
		},
		{
			"id": 2,
			"title": "1984",
			"author": "Джордж Оруэлл",
			"genre": "Антиутопия",
		},
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}