package mock

import (
	"FreeLib/internal/models"
	"errors"
)

type MockUserRepository struct {
	users []models.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: []models.User{
			{
				ID: 1,
				Username: "test1",
				Email: "test1@gmail.com",
				IsAdmin: false,
				PasswordHash: "12345678",
			},
		},
	}
}

func (m *MockUserRepository) CreateUser(user *models.User) error {
	for _, u := range m.users {
		if u.Email == user.Email {
			return errors.New("user with such an email already exists")
		}
		if u.Username == user.Username {
			return errors.New("user with such an username already exists")
		} 
	}

	maxID := uint(0)
	for _, u := range m.users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}
	user.ID = maxID + 1 

	m.users = append(m.users, *user)

	return nil
}