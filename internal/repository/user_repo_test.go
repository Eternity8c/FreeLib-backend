package repository_test

import (
	"FreeLib/internal/models"
	"FreeLib/internal/repository/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateMockUser() models.User {
	return models.User{
		ID:           1,
		Username:     "1",
		Email:        "1",
		IsAdmin:      false,
		PasswordHash: "1",
	}
}

func TestRegister(t *testing.T) {
	mockUser := CreateMockUser()
	mockRepo := mock.NewMockUserRepository()

	err := mockRepo.CreateUser(&mockUser)
	assert.Nil(t, err)

	err = mockRepo.CreateUser(&mockUser)
	assert.NotNil(t, err)
}

func TestAuntificateUser(t *testing.T) {
	mockUser := CreateMockUser()
	mockRepo := mock.NewMockUserRepository()

	lrTrue := models.LoginRequest{
		Email:    "1",
		Password: "1",
	}

	mockRepo.CreateUser(&mockUser)

	resUserTrue, err := mockRepo.AuntificationUser(&lrTrue)

	assert.Nil(t, err)
	assert.Equal(t, *resUserTrue, mockRepo.Users[0])

	lrFalse := models.LoginRequest{
		Email:    "1221212",
		Password: "1",
	}

	resUserFalse, err := mockRepo.AuntificationUser(&lrFalse)

	assert.Nil(t, resUserFalse)
	assert.NotNil(t, err)
}
