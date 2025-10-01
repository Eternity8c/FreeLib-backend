package repository

import (
	"FreeLib/internal/models"
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

