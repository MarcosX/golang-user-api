package domain

import (
	"github.com/brizenox/golang-user-api/internal/db"
)

type UserRepository interface {
	GetUser(id string) (*db.User, error)
	GetUserByEmail(name string) (*db.User, error)
	CreateUser(name string, email string, password string) (*db.User, error)
	UpdateUser(id string, name string, email string, password string) (*db.User, error)
}

type realUserRepository struct {
}

func (u *realUserRepository) GetUser(id string) (*db.User, error) {
	user := db.GetUserById(id)
	if user == nil {
		return nil, &db.ErrUserNotFound{Id: id}
	}
	return user, nil
}

func (u *realUserRepository) GetUserByEmail(email string) (*db.User, error) {
	user := db.GetUserByEmail(email)
	if user == nil {
		return nil, &db.ErrUserNotFound{Id: email}
	}
	return user, nil
}

func (u *realUserRepository) UpdateUser(id string, name string, email string, password string) (*db.User, error) {
	user, err := u.GetUser(id)
	if err != nil {
		return nil, err
	}
	user.Name = name
	user.Email = email
	user.Password = password
	db.SaveUser(user)
	return user, nil
}

func (u *realUserRepository) CreateUser(name string, email string, password string) (*db.User, error) {
	return db.CreateUser(name, email, password)
}

func NewUserRepository() UserRepository {
	return &realUserRepository{}
}
