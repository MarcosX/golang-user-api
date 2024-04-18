package domain

import (
	"github.com/brizenox/golang-user-api/internal/db"
)

type UserRepository interface {
	GetUser(id string) (*db.User, error)
	GetUserByEmail(name string) (*db.User, error)
	CreateUser(name string, email string, password string) (*db.User, error)
}

type realUserRepository struct {
}

func (u *realUserRepository) GetUser(id string) (*db.User, error) {
	for _, v := range db.GetAllUsers() {
		if v.Id == id {
			return v, nil
		}
	}
	return nil, &db.ErrUserNotFound{Id: id}
}

func (u *realUserRepository) GetUserByEmail(email string) (*db.User, error) {
	for _, v := range db.GetAllUsers() {
		if v.Email == email {
			return v, nil
		}
	}
	return nil, &db.ErrUserNotFound{Id: email}
}

func (u *realUserRepository) CreateUser(name string, email string, password string) (*db.User, error) {
	return db.CreateUser(name, email, password)
}

func NewUserRepository() UserRepository {
	return &realUserRepository{}
}