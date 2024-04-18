package domain

import (
	"fmt"

	"github.com/brizenox/golang-user-api/internal/db"
)

var UserSample = db.NewUser("0", "User", "user@email.com", "pass")

type mockUserRepository struct {
	lastId int
	users  map[string]*db.User
}

func (u *mockUserRepository) GetUser(id string) (*db.User, error) {
	elem, ok := u.users[id]
	if ok {
		return elem, nil
	}
	return nil, &db.ErrUserNotFound{Id: id}
}

func (u *mockUserRepository) GetUserByEmail(email string) (*db.User, error) {
	for _, v := range u.users {
		if v.Email == email {
			return v, nil
		}
	}
	return nil, &db.ErrUserNotFound{Id: email}
}

func (u *mockUserRepository) CreateUser(name string, email string, password string) (*db.User, error) {
	if name == "Crash()" {
		return nil, fmt.Errorf("Crash() called")
	}
	u.lastId += 1
	user := db.NewUser(fmt.Sprintf("%d", u.lastId), name, email, password)
	u.users[user.Id] = user
	return user, nil
}

func (u *mockUserRepository) UpdateUser(id string, name string, email string, password string) (*db.User, error) {
	if name == "Crash()" {
		return nil, fmt.Errorf("Crash() called")
	}
	user := db.NewUser(id, name, email, password)
	u.users[id] = user
	return user, nil
}

func NewMockUserRepository() *mockUserRepository {
	return &mockUserRepository{
		lastId: 0,
		users: map[string]*db.User{
			"0": UserSample,
		},
	}
}
