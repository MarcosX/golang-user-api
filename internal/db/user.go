package db

import "github.com/google/uuid"

var db map[string]*User

func init() {
	db = make(map[string]*User)
	db["0"] = &User{
		Id:       "0",
		Name:     "User",
		Email:    "user@email.com",
		Password: "pass",
	}
}

type User struct {
	Id       string `json:"-"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (u *User) PasswordMatches(password string) bool {
	return u.Password == password
}

func CreateUser(name string, email string, password string) (*User, error) {
	id := uuid.Must(uuid.NewRandom()).String()
	user := &User{
		Id:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}
	db[id] = user
	return user, nil
}

func GetAllUsers() []*User {
	users := make([]*User, 0, len(db))
	for _, user := range db {
		users = append(users, user)
	}
	return users
}
