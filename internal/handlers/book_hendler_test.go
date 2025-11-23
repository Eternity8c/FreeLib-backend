package handlers

import (
	"FreeLib/internal/models"
	"FreeLib/internal/repository/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {
	bookRepo := mock.NewMockBookRepo()
	book := models.Book{
		ID: 1,
		Author: "Булгаков М.А.",
		Title: "Мастер и Маргарита",
		Description: "1234567890",
		Content: "111",
		CoverURL: "222",
		CreatedAt: time.Now(),
	}
	
	err := bookRepo.Create(&book)
	assert.Nil(t, err)
}