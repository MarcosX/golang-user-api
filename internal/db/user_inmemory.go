package db

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

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

func NewUser(id, name, email, password string) *User {
	return &User{
		Id:       id,
		Name:     name,
		Email:    email,
		Password: hashAndSaltPassword(password),
	}
}

func (u *User) PasswordMatches(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func CreateUser(name string, email string, password string) (*User, error) {
	id := uuid.Must(uuid.NewRandom()).String()
	user := NewUser(id, name, email, password)
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

func SaveUser(user *User) {
	user.Password = hashAndSaltPassword(user.Password)
	db[user.Id] = user
}

func hashAndSaltPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
