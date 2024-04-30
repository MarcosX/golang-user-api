package domain

import (
	"fmt"

	"github.com/brizenox/golang-user-api/internal/db"
)

var UserSample = db.NewUser("User", "user@email.com", "pass")

type mockUserRepository struct {
	users []*db.User
}

func (repository *mockUserRepository) GetUser(id string) (*db.User, error) {
	var user *db.User
	for _, u := range repository.users {
		if u.Id == id {
			user = u
			break
		}
	}
	if user != nil {
		return user, nil
	}
	return nil, &db.ErrUserNotFound{Id: id}
}

func (repository *mockUserRepository) GetUserByEmail(email string) (*db.User, error) {
	var user *db.User
	for _, u := range repository.users {
		if u.Email == email {
			user = u
			break
		}
	}
	if user != nil {
		return user, nil
	}
	return nil, &db.ErrUserNotFound{Id: email}
}

func (repository *mockUserRepository) CreateUser(name string, email string, password string) (*db.User, error) {
	if name == "Crash()" {
		return nil, fmt.Errorf("Crash() called")
	}
	user := db.NewUser(name, email, password)
	repository.users = append(repository.users, user)
	return user, nil
}

func (repository *mockUserRepository) UpdateUser(id string, name string, email string, password string) (*db.User, error) {
	if name == "Crash()" {
		return nil, fmt.Errorf("Crash() called")
	}
	user := db.NewUser(name, email, password)
	repository.users = append(repository.users, user)
	return user, nil
}

func NewMockUserRepository() *mockUserRepository {
	return &mockUserRepository{
		users: []*db.User{UserSample},
	}
}
