package mock

import (
	"FreeLib/internal/models"
	"FreeLib/internal/repository"
	"errors"
	"time"
)

type MockBookRepo struct {
	Books []models.Book
}

func NewMockBookRepo() repository.BookRepository {
	return &MockBookRepo{
		Books: []models.Book{},
	}
}

func (m *MockBookRepo) GetAll() ([]models.Book, error) {
	return m.Books, nil
}

func (m *MockBookRepo) GetByID(id uint) (*models.Book, error) {
	for i := 0; i < len(m.Books); i++ {
		if m.Books[i].ID == id {
			return &m.Books[i], nil
		}
	}
	return nil, errors.New("ID not fount")
}

func (m *MockBookRepo) Update(book *models.Book) error {
	for i, mockBook := range m.Books {
		if mockBook.ID == book.ID {
			m.Books[i] = *book
			return nil
		}
	}
	return errors.New("book not found")
}

func (m *MockBookRepo) Create(book *models.Book) error {
	maxID := uint(0)
	for _, b := range m.Books {
		if b.ID > maxID {
			maxID = b.ID
		}
	}
	for i := 0; i < len(m.Books); i++ {
		if m.Books[i].Author == book.Author &&
			m.Books[i].Title == book.Title {
			return errors.New("book already exists")
		}
	}
	book.ID = maxID + 1
	book.CreatedAt = time.Now()
	m.Books = append(m.Books, *book)
	return nil
}

func (m *MockBookRepo) Delete(id uint) error {
	for i := 0; i < len(m.Books); i++ {
		if id == m.Books[i].ID {
			m.Books = append(m.Books[:i], m.Books[i+1:]...)
			return nil
		}
	}
	return errors.New("ID not found")
}

func (m *MockBookRepo) AddFavorite(userID uint, bookID uint) error {
	if userID == 0 {
		return errors.New("userID cannot be zero")
	}

	for _, book := range m.Books {
		if book.ID == bookID {
			return nil
		}
	}

	return errors.New("bookID not found")
}

func (m *MockBookRepo) DeleteFavorite(userID uint, bookID uint) error {
	return nil
}

func (m *MockBookRepo) GetAllFavorite(userID uint) ([]models.Book, error) {
	return nil, nil
}
