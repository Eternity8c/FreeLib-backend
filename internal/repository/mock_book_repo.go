package repository

import (
	"FreeLib/internal/models"
	"errors"
	"time"
)

type MockBookRepo struct {}

var mockBooks = []models.Book {
	{
		ID: 1,
		Title: "Мастер и Маргарита",
		Author: "M.А. Булгаков",
		Description: "Роман о дьяволе, посетившем Москву 1930-х годов",
		CreatedAt: time.Now(),
	},
		{
		ID: 2,
		Title: "Преступление и наказание",
		Author: "Ф.М. Достоевский",
		Description: "Психологическая драма о студенте, совершившем убийство",
		CreatedAt: time.Now(),
	},
	{
		ID: 3,
		Title: "Война и мир",
		Author: "Л.Н. Толстой",
		Description: "Эпопея о жизни русского общества во время войны с Наполеоном",
		CreatedAt: time.Now(),
	},
	{
		ID: 4,
		Title: "Отцы и дети",
		Author: "И.С. Тургенев",
		Description: "Роман о конфликте поколений и нигилизме",
		CreatedAt: time.Now(),
	},
	{
		ID: 5,
		Title: "Мёртвые души",
		Author: "Н.В. Гоголь",
		Description: "Поэма о афере с покупкой умерших крестьян",
		CreatedAt: time.Now(),
	},
	{
		ID: 6,
		Title: "Анна Каренина",
		Author: "Л.Н. Толстой",
		Description: "Трагическая история любви замужней женщины",
		CreatedAt: time.Now(),
	},
}

func (m *MockBookRepo) GetAll() ([]models.Book, error) {
	return mockBooks ,nil
}

func (m *MockBookRepo) GetByID(id uint) (*models.Book, error) {
	for i := 0; i < len(mockBooks); i++ {
		if mockBooks[i].ID == id {
			return &mockBooks[i], nil
		}
	}
	return nil, errors.New("Не найден нужный ID")
}

func (m *MockBookRepo) Search(query string) ([]models.Book, error) {
	var books []models.Book
	for i := 0; i < len(mockBooks); i++ {
		if mockBooks[i].Author == query || mockBooks[i].Title == query {
			books = append(books, mockBooks[i])
		}
	}

	if len(books) == 0 {
		return nil, errors.New("Не удалось найти книгу по заданным параметрам")
	}

	return books, nil
}


func (m *MockBookRepo) Create(book *models.Book) error {
	maxID := uint(0);
	for _, b := range mockBooks {
		if b.ID > maxID {
			maxID = b.ID
		}
	}
	for i := 0; i < len(mockBooks); i++ {
		if mockBooks[i].Author == book.Author && 
		mockBooks[i].Title == book.Title {
			return errors.New("Такая книга уже есть")
		}
	}
	book.ID = maxID + 1
	book.CreatedAt = time.Now()
	mockBooks = append(mockBooks, *book)
	return nil
}

func (m *MockBookRepo) Delete(id uint) error {
	for i := 0; i < len(mockBooks); i++ {
		if id == mockBooks[i].ID {
			mockBooks = append(mockBooks[:i], mockBooks[i+1:]...)
			return nil
		}
	}
	return errors.New("ID не найден")
}