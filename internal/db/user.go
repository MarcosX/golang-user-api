package db

type User struct {
	Id    string `json:"-"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository interface {
	GetUser(id string) (*User, error)
}

type realUserRepository struct {
}

func (u *realUserRepository) GetUser(id string) (*User, error) {
	panic("Not implemented")
}

func NewUserRepository() UserRepository {
	return &realUserRepository{}
}
