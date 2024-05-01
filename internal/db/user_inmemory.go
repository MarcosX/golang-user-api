package db

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type InMemoryDb struct {
	usersById    map[string]*User
	usersByEmail map[string]*User
}

func NewInMemoryDb() *InMemoryDb {
	return &InMemoryDb{
		usersById:    make(map[string]*User),
		usersByEmail: make(map[string]*User),
	}
}

type User struct {
	Id       string `json:"-"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	u := &User{
		Id:       uuid.Must(uuid.NewRandom()).String(),
		Name:     name,
		Email:    email,
		Password: password,
	}
	hashedPassword, err := hashAndSaltPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.Password = hashedPassword
	return u, nil
}

func (u *User) PasswordMatches(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (db *InMemoryDb) CreateUser(name string, email string, password string) (*User, error) {
	if db.usersByEmail[email] != nil {
		return nil, &ErrUserAlreadyExists{Id: email}
	}
	user, err := NewUser(name, email, password)
	if err != nil {
		return nil, err
	}
	db.usersById[user.Id] = user
	db.usersByEmail[user.Email] = user
	return user, nil
}

func (db *InMemoryDb) GetAllUsers() []*User {
	users := make([]*User, 0, len(db.usersById))
	for _, user := range db.usersById {
		users = append(users, user)
	}
	return users
}

func (db *InMemoryDb) GetUserById(id string) *User {
	return db.usersById[id]
}

func (db *InMemoryDb) GetUserByEmail(email string) *User {
	return db.usersByEmail[email]
}

func (db *InMemoryDb) SaveUser(u *User) error {
	if db.usersByEmail[u.Email] != nil {
		return &ErrUserAlreadyExists{Id: u.Email}
	}
	hashedPassword, err := hashAndSaltPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	db.usersById[u.Id] = u
	db.usersByEmail[u.Email] = u
	return nil
}

func hashAndSaltPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hashedPassword), err
}
