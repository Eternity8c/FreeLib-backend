package repository

import "FreeLib/internal/models"

type BookReposytory interface {
	GetAll() ([]models.Book, error)
	GetByID(id uint) (*models.Book, error)
	Create(book *models.Book) error
	Search(query string) ([]models.Book, error)
}