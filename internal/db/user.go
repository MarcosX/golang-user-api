package db

import "fmt"

type User struct {
	Id    string
	Name  string
	Email string
}

type UserRepository interface {
	GetUser(id string) (*User, error)
}

type realUserRepository struct {
	users map[string]*User
}

func (u *realUserRepository) GetUser(id string) (*User, error) {
	panic("Not implemented")
}

func NewUserRepository() UserRepository {
	return &realUserRepository{
		users: map[string]*User{
			"0": {
				Email: "user@email.com",
				Name:  "User",
				Id:    "0",
			},
		},
	}
}

type ErrUserNotFound struct {
	Id string
}

func (e *ErrUserNotFound) Error() string {
	return fmt.Sprintf("could not find user %s", e.Id)
}
