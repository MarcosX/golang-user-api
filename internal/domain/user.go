package domain

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
	return &User{
		Id:    id,
		Name:  "User",
		Email: "user@email.com",
	}, nil
}

func NewUserRepository() UserRepository {
	return &realUserRepository{}
}
