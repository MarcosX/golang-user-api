package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPassword(t *testing.T) {
	user := User{
		Password: "admin123",
	}
	assert.True(t, user.PasswordMatches("admin123"))
}

func TestWrongPassword(t *testing.T) {
	user := User{
		Password: "admin123",
	}
	assert.False(t, user.PasswordMatches("admin321"))
}

func TestInitInMemoryDB(t *testing.T) {
	assert.GreaterOrEqual(t, len(db), 1)
}

func TestCreateUser(t *testing.T) {
	user, err := CreateUser("User", "user@email.com", "pass")
	if assert.NoError(t, err) {
		assert.NotNil(t, user)
		assert.NotEmpty(t, user.Id)
		assert.Equal(t, "User", user.Name)
		assert.Equal(t, "user@email.com", user.Email)
		assert.Equal(t, "pass", user.Password)
	}
	assert.GreaterOrEqual(t, len(db), 2)
}

func TestSaveUser(t *testing.T) {
	user, _ := CreateUser("User", "user@email.com", "pass")
	user.Name = "User Updated"
	user.Email = "another.email@email.com"
	user.Password = "pass123"
	SaveUser(user)
	user = db[user.Id]

	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Id)
	assert.Equal(t, "User Updated", user.Name)
	assert.Equal(t, "another.email@email.com", user.Email)
	assert.Equal(t, "pass123", user.Password)
}

func TestGetAllUsers(t *testing.T) {
	user, _ := CreateUser("User", "user@email.com", "pass")
	users := GetAllUsers()
	assert.Contains(t, users, user)
}
