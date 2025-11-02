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
		users: []models.User{},
	}
}

func (m *MockUserRepository) CreateUser(user *models.User) error {
	for _, u := range m.users {
		if u.Email == user.Email {
			return errors.New("User с таким email существует")
		}
		if u.Username == user.Username {
			return errors.New("User с таким именем существует")
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