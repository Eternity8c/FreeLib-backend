package mock

import (
	"FreeLib/internal/models"
	"errors"
)

type MockUserRepository struct {
	Users []models.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		Users: []models.User{},
	}
}

func (m *MockUserRepository) CreateUser(user *models.User) error {
	for _, u := range m.Users {
		if u.Email == user.Email {
			return errors.New("user with such an email already exists")
		}
		if u.Username == user.Username {
			return errors.New("user with such an username already exists")
		}
	}

	maxID := uint(0)
	for _, u := range m.Users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}
	user.ID = maxID + 1

	m.Users = append(m.Users, *user)

	return nil
}

func (m *MockUserRepository) AuntificationUser(lr *models.LoginRequest) (*models.User, error) {
	ID := 0
	for i := 0; i < len(m.Users); i++ {
		if m.Users[i].Email != lr.Email {
			return nil, errors.New("emails not match")
		}
		if m.Users[i].PasswordHash != lr.Password {
			return nil, errors.New("passwords not match")
		} else {
			ID = i
		}
	}
	return &m.Users[ID], nil
}
