package repository_test

import (
	"FreeLib/internal/models"
	"FreeLib/internal/repository"
	"FreeLib/internal/repository/mock"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetByIDAndCreate(t *testing.T) {
	mockRepo := mock.NewMockBookRepo()
	mockBook := models.Book{
		Author:      "1",
		Title:       "1",
		Description: "1",
		Content:     "1",
		CoverURL:    "1",
		Genre:       "1",
	}

	err := mockRepo.Create(&mockBook)
	assert.Nil(t, err)

	err = mockRepo.Create(&mockBook)
	assert.NotNil(t, err)

	bookByIDTrue, err := mockRepo.GetByID(uint(1))

	assert.NotNil(t, bookByIDTrue)
	assert.Equal(t, mockBook, *bookByIDTrue)
	assert.Nil(t, err)

	bookByIDFalse, err := mockRepo.GetByID(uint(333))

	assert.Nil(t, bookByIDFalse)
	assert.Error(t, err)
}

func SeedMockRepository(repo repository.BookRepository) {
	for i := 0; i < 10; i++ {
		iStr := strconv.Itoa(i)
		mockBook := models.Book{
			Author:      iStr,
			Title:       "1",
			Description: "1",
			Content:     "1",
			CoverURL:    "1",
			Genre:       "1",
		}
		repo.Create(&mockBook)
	}
}

func TestGetAll(t *testing.T) {
	mockRepo1 := mock.NewMockBookRepo()
	SeedMockRepository(mockRepo1)

	mockRepo2 := mock.NewMockBookRepo()
	SeedMockRepository(mockRepo2)

	res1, err := mockRepo1.GetAll()
	res2, err := mockRepo1.GetAll()

	assert.NotNil(t, mockRepo1)
	assert.NotNil(t, mockRepo2)
	assert.Equal(t, res1, res2)
	assert.Nil(t, err)
}

func TestSearch(t *testing.T) {
	mockRepo := mock.NewMockBookRepo()
	SeedMockRepository(mockRepo)
	mockBookTrue := models.Book{
		ID:          1,
		Author:      "0",
		Title:       "1",
		Description: "1",
		Content:     "1",
		CoverURL:    "1",
		Genre:       "1",
		CreatedAt:   time.Now(),
	}

	bookByID, err := mockRepo.GetByID(uint(1))

	assert.NotNil(t, bookByID)
	assert.Equal(t, mockBookTrue, *bookByID)
	assert.Nil(t, err)

	err = mockRepo.Update(&mockBookTrue)
	assert.Nil(t, err)
	assert.Equal(t, mockBookTrue, *bookByID)

	mockBookFalse := models.Book{
		ID:          111,
		Author:      "1111",
		Title:       "1",
		Description: "1",
		Content:     "1",
		CoverURL:    "1",
		Genre:       "1",
		CreatedAt:   time.Now(),
	}

	err = mockRepo.Update(&mockBookFalse)
	assert.NotNil(t, err)
}
