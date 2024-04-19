package db

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type db struct {
	usersById    map[string]*User
	usersByEmail map[string]*User
}

var usersDb *db

func init() {
	usersDb = &db{
		usersById:    make(map[string]*User),
		usersByEmail: make(map[string]*User),
	}
	sampleUser := &User{
		Id:       "0",
		Name:     "User",
		Email:    "user@email.com",
		Password: "pass",
	}
	usersDb.usersById[sampleUser.Id] = sampleUser
	usersDb.usersByEmail[sampleUser.Email] = sampleUser
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
	if usersDb.usersByEmail[email] != nil {
		return nil, &ErrUserAlreadyExists{Id: email}
	}
	id := uuid.Must(uuid.NewRandom()).String()
	user := NewUser(id, name, email, password)
	usersDb.usersById[user.Id] = user
	usersDb.usersByEmail[user.Email] = user
	return user, nil
}

func GetAllUsers() []*User {
	users := make([]*User, 0, len(usersDb.usersById))
	for _, user := range usersDb.usersById {
		users = append(users, user)
	}
	return users
}

func GetUserById(id string) *User {
	return usersDb.usersById[id]
}

func GetUserByEmail(email string) *User {
	return usersDb.usersByEmail[email]
}

func SaveUser(user *User) error {
	if usersDb.usersByEmail[user.Email] != nil {
		return &ErrUserAlreadyExists{Id: user.Email}
	}
	user.Password = hashAndSaltPassword(user.Password)
	usersDb.usersById[user.Id] = user
	usersDb.usersByEmail[user.Email] = user
	return nil
}

func hashAndSaltPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hashedPassword)
}
