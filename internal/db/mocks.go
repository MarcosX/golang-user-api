package db

import "github.com/brizenox/golang-user-api/internal/domain"

var UserSample = &domain.User{
	Email: "user@email.com",
	Name:  "User",
	Id:    "0",
}

type MockUserRepository struct {
	users map[string]*domain.User
}

func (u *MockUserRepository) GetUser(id string) (*domain.User, error) {
	elem, ok := u.users[id]
	if ok {
		return elem, nil
	}
	return nil, &ErrUserNotFound{Id: id}
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: map[string]*domain.User{
			"0": UserSample,
		},
	}
}
