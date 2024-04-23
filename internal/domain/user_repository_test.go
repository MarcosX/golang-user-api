package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserWithEmptyName(t *testing.T) {
	u := NewUserRepository()
	_, err := u.CreateUser("", "user@email", "pass123")
	assert.Error(t, err)
}

func TestCreateUserWithEmptyEmail(t *testing.T) {
	u := NewUserRepository()
	_, err := u.CreateUser("User", "", "pass123")
	assert.Error(t, err)
}

func TestCreateUserWithInvalidEmail(t *testing.T) {
	u := NewUserRepository()
	_, err := u.CreateUser("User", "invalid-email", "pass123")
	assert.Error(t, err)
}

func TestCreateUserWithEmptyPassword(t *testing.T) {
	u := NewUserRepository()
	_, err := u.CreateUser("User", "user@email.com", "")
	assert.Error(t, err)
}

func TestSaveUserWithEmptyName(t *testing.T) {
	u := NewUserRepository()
	user, _ := u.CreateUser("User", "user@email.com", "pass123")
	_, err := u.UpdateUser(user.Id, "", user.Email, user.Password)
	assert.Error(t, err)
}

func TestSaveUserWithEmptyEmail(t *testing.T) {
	u := NewUserRepository()
	user, _ := u.CreateUser("User", "user@email.com", "pass123")
	_, err := u.UpdateUser(user.Id, user.Name, "", user.Password)
	assert.Error(t, err)
}

func TestSaveUserWithInvalidEmail(t *testing.T) {
	u := NewUserRepository()
	user, _ := u.CreateUser("User", "user@email.com", "pass123")
	_, err := u.UpdateUser(user.Id, user.Name, "invalid-email", user.Password)
	assert.Error(t, err)
}

func TestSaveUserWithEmptyPassword(t *testing.T) {
	u := NewUserRepository()
	user, _ := u.CreateUser("User", "user@email.com", "pass123")
	_, err := u.UpdateUser(user.Id, user.Name, user.Email, "")
	assert.Error(t, err)
}
