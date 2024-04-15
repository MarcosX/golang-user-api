package domain

type User struct {
	Id       string `json:"-"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (u *User) PasswordMatches(password string) bool {
	return u.Password == password
}

type UserRepository interface {
	GetUser(id string) (*User, error)
	GetUserByEmail(name string) (*User, error)
}

type realUserRepository struct {
}

func (u *realUserRepository) GetUser(id string) (*User, error) {
	return &User{
		Id:       id,
		Name:     "User",
		Email:    "user@email.com",
		Password: "pass",
	}, nil
}

func (u *realUserRepository) GetUserByEmail(email string) (*User, error) {
	return &User{
		Id:       "0",
		Name:     "User",
		Email:    "user@email.com",
		Password: "pass",
	}, nil
}

func NewUserRepository() UserRepository {
	return &realUserRepository{}
}
