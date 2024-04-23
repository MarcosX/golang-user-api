package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPassword(t *testing.T) {
	db := NewInMemoryDb()
	user, _ := db.CreateUser("User", "email@email.com", "admin123")
	assert.True(t, user.PasswordMatches("admin123"))
}

func TestWrongPassword(t *testing.T) {
	db := NewInMemoryDb()
	user, _ := db.CreateUser("User", "email@email.com", "admin123")
	assert.False(t, user.PasswordMatches("admin321"))
}

func TestCreateUser(t *testing.T) {
	db := NewInMemoryDb()
	user, err := db.CreateUser("User", "user@email.com", "pass")
	if assert.NoError(t, err) {
		assert.NotNil(t, user)
		assert.NotEmpty(t, user.Id)
		assert.Equal(t, "User", user.Name)
		assert.Equal(t, "user@email.com", user.Email)
		assert.True(t, user.PasswordMatches("pass"))
	}
	assert.Equal(t, len(db.usersById), 1)
	assert.Equal(t, len(db.usersByEmail), 1)
}

func TestSaveUser(t *testing.T) {
	db := NewInMemoryDb()
	user, _ := db.CreateUser("User", "user@email.com", "pass")
	user.Name = "User Updated"
	user.Email = "another.email@email.com"
	user.Password = "pass123"
	db.SaveUser(user)
	user = db.usersById[user.Id]

	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Id)
	assert.Equal(t, "User Updated", user.Name)
	assert.Equal(t, "another.email@email.com", user.Email)
	assert.True(t, user.PasswordMatches("pass123"))
}

func TestGetAllUsers(t *testing.T) {
	db := NewInMemoryDb()
	user, _ := db.CreateUser("User", "user@email.com", "pass")
	users := db.GetAllUsers()
	assert.Contains(t, users, user)
}

func TestGetUserById(t *testing.T) {
	db := NewInMemoryDb()
	user, _ := db.CreateUser("User", "user@email.com", "pass")

	userFound := db.GetUserById(user.Id)
	assert.NotNil(t, userFound)

}
func TestGetUserByEmail(t *testing.T) {
	db := NewInMemoryDb()
	user, _ := db.CreateUser("User", "user@email.com", "pass")

	userFound := db.GetUserByEmail(user.Email)
	assert.NotNil(t, userFound)
}

func TestCreateUserWithDuplicateEmail(t *testing.T) {
	db := NewInMemoryDb()
	db.CreateUser("User", "user@email.com", "pass")

	_, err := db.CreateUser("User", "user@email.com", "pass")
	assert.Error(t, err)
}

func TestSaveUserWithDuplicateEmail(t *testing.T) {
	db := NewInMemoryDb()
	db.CreateUser("User", "user@email.com", "pass")
	secondUser, _ := db.CreateUser("User 2", "anotheruser@email.com", "pass")
	secondUser.Email = "user@email.com"

	err := db.SaveUser(secondUser)
	assert.Error(t, err)
}
