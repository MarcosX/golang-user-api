package domain

import "github.com/brizenox/golang-user-api/internal/db"

var UserSample = &User{
	Email:    "user@email.com",
	Name:     "User",
	Id:       "0",
	Password: "pass",
}

type mockUserRepository struct {
	users map[string]*User
}

func (u *mockUserRepository) GetUser(id string) (*User, error) {
	elem, ok := u.users[id]
	if ok {
		return elem, nil
	}
	return nil, &db.ErrUserNotFound{Id: id}
}

func (u *mockUserRepository) GetUserByEmail(email string) (*User, error) {
	for _, v := range u.users {
		if v.Email == email {
			return v, nil
		}
	}
	return nil, &db.ErrUserNotFound{Id: email}
}

func NewMockUserRepository() *mockUserRepository {
	return &mockUserRepository{
		users: map[string]*User{
			"0": UserSample,
		},
	}
}
