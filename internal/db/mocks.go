package db

var UserSample = &User{
	Email: "user@email.com",
	Name:  "User",
	Id:    "0",
}

type MockUserRepository struct {
	users map[string]*User
}

func (u *MockUserRepository) GetUser(id string) (*User, error) {
	elem, ok := u.users[id]
	if ok {
		return elem, nil
	}
	return nil, &ErrUserNotFound{Id: id}
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: map[string]*User{
			"0": UserSample,
		},
	}
}
