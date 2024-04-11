package db

import "errors"

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

type MockSessionRepository struct {
	sessions map[string]*Session
}

func (s *MockSessionRepository) GetSession(sessionId string) (*Session, error) {
	elem, ok := s.sessions[sessionId]
	if ok {
		return elem, nil
	}
	return nil, errors.New("session not found")
}

func NewMockSessionRepository() *MockSessionRepository {
	return &MockSessionRepository{
		sessions: map[string]*Session{
			"samplesession": {
				SignedToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIwIn0.hR9oaCu7Ud_Mr-QENEc-K6DLdZBaReap1rpvgnyEPU0",
			},
			"wrongsecretsession": {
				SignedToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIwIn0.nL3zQibYzBCqwzILJ6KJQSiYEEXjxqnu5rM0_U-ZH0E",
			},
		},
	}
}
