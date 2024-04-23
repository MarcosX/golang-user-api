package domain

import (
	"regexp"

	"github.com/brizenox/golang-user-api/internal/db"
)

type UserRepository interface {
	GetUser(id string) (*db.User, error)
	GetUserByEmail(name string) (*db.User, error)
	CreateUser(name string, email string, password string) (*db.User, error)
	UpdateUser(id string, name string, email string, password string) (*db.User, error)
}

type realUserRepository struct {
	db *db.InMemoryDb
}

func NewUserRepository() UserRepository {
	return &realUserRepository{
		db: db.NewInMemoryDb(),
	}
}

func (u *realUserRepository) GetUser(id string) (*db.User, error) {
	user := u.db.GetUserById(id)
	if user == nil {
		return nil, &db.ErrUserNotFound{Id: id}
	}
	return user, nil
}

func (u *realUserRepository) GetUserByEmail(email string) (*db.User, error) {
	user := u.db.GetUserByEmail(email)
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
	if err := validate(user); err != nil {
		return nil, err
	}
	u.db.SaveUser(user)
	return user, nil
}

func (u *realUserRepository) CreateUser(name string, email string, password string) (*db.User, error) {
	user := db.NewUser(name, email, password)
	if err := validate(user); err != nil {
		return nil, err
	}
	return u.db.CreateUser(name, email, password)
}

func validate(u *db.User) error {
	if u.Name == "" {
		return &db.ErrValidationFailed{Field: "name"}
	}
	if u.Email == "" {
		return &db.ErrValidationFailed{Field: "email"}
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.Email) {
		return &db.ErrValidationFailed{Field: "email"}
	}
	if u.Password == "" || u.PasswordMatches("") {
		return &db.ErrValidationFailed{Field: "password"}
	}
	return nil
}
