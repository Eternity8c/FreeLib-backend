package handlers

import (
	"FreeLib/internal/models"
	"FreeLib/internal/repository/mock"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var userRepo = mock.NewMockUserRepository()

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req models.RegisteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Password hashing error:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	user := models.User {
		Username: req.Username,
		Email: req.Email,
		PasswordHash: string(passwordHash),
	}

	err = userRepo.CreateUser(&user)
	if err != nil {
		log.Println("Create user error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{} {
		"id": user.ID,
		"username": user.Username,
		"email": user.Email,
		"message": "User registred successfully",
	})
}