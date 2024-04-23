package db

import (
	"regexp"

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

func NewUser(id, name, email, password string) (*User, error) {
	u := &User{
		Id:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}
	if err := u.Validate(); err != nil {
		return nil, err
	}
	u.Password = hashAndSaltPassword(u.Password)
	return u, nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return &ErrValidationFailed{Field: "name"}
	}
	if u.Email == "" {
		return &ErrValidationFailed{Field: "email"}
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.Email) {
		return &ErrValidationFailed{Field: "email"}
	}
	if u.Password == "" {
		return &ErrValidationFailed{Field: "password"}
	}
	return nil
}

func (u *User) PasswordMatches(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func CreateUser(name string, email string, password string) (*User, error) {
	if usersDb.usersByEmail[email] != nil {
		return nil, &ErrUserAlreadyExists{Id: email}
	}
	id := uuid.Must(uuid.NewRandom()).String()
	user, err := NewUser(id, name, email, password)
	if err != nil {
		return nil, err
	}
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

func (u *User) SaveUser() error {
	if err := u.Validate(); err != nil {
		return err
	}
	if usersDb.usersByEmail[u.Email] != nil {
		return &ErrUserAlreadyExists{Id: u.Email}
	}
	u.Password = hashAndSaltPassword(u.Password)
	usersDb.usersById[u.Id] = u
	usersDb.usersByEmail[u.Email] = u
	return nil
}

func hashAndSaltPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hashedPassword)
}
