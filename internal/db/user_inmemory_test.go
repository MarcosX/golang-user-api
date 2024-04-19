package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPassword(t *testing.T) {
	defer clearUsersDb()
	user, _ := CreateUser("User", "email@email.com", "admin123")
	assert.True(t, user.PasswordMatches("admin123"))
}

func TestWrongPassword(t *testing.T) {
	defer clearUsersDb()
	user, _ := CreateUser("User", "email@email.com", "admin123")
	assert.False(t, user.PasswordMatches("admin321"))
}

func TestCreateUser(t *testing.T) {
	defer clearUsersDb()
	user, err := CreateUser("User", "user@email.com", "pass")
	if assert.NoError(t, err) {
		assert.NotNil(t, user)
		assert.NotEmpty(t, user.Id)
		assert.Equal(t, "User", user.Name)
		assert.Equal(t, "user@email.com", user.Email)
		assert.True(t, user.PasswordMatches("pass"))
	}
	assert.Equal(t, len(usersDb.usersById), 1)
	assert.Equal(t, len(usersDb.usersByEmail), 1)
}

func TestSaveUser(t *testing.T) {
	defer clearUsersDb()
	user, _ := CreateUser("User", "user@email.com", "pass")
	user.Name = "User Updated"
	user.Email = "another.email@email.com"
	user.Password = "pass123"
	SaveUser(user)
	user = usersDb.usersById[user.Id]

	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Id)
	assert.Equal(t, "User Updated", user.Name)
	assert.Equal(t, "another.email@email.com", user.Email)
	assert.True(t, user.PasswordMatches("pass123"))
}

func TestGetAllUsers(t *testing.T) {
	defer clearUsersDb()
	user, _ := CreateUser("User", "user@email.com", "pass")
	users := GetAllUsers()
	assert.Contains(t, users, user)
}

func TestGetUserById(t *testing.T) {
	defer clearUsersDb()
	user, _ := CreateUser("User", "user@email.com", "pass")

	userFound := GetUserById(user.Id)
	assert.NotNil(t, userFound)

}
func TestGetUserByEmail(t *testing.T) {
	defer clearUsersDb()
	user, _ := CreateUser("User", "user@email.com", "pass")

	userFound := GetUserByEmail(user.Email)
	assert.NotNil(t, userFound)
}

func TestCreateUserWithDuplicateEmail(t *testing.T) {
	defer clearUsersDb()
	CreateUser("User", "user@email.com", "pass")

	_, err := CreateUser("User", "user@email.com", "pass")
	assert.Error(t, err)
}

func TestSaveUserWithDuplicateEmail(t *testing.T) {
	defer clearUsersDb()
	CreateUser("User", "user@email.com", "pass")
	secondUser, _ := CreateUser("User 2", "anotheruser@email.com", "pass")
	secondUser.Email = "user@email.com"

	err := SaveUser(secondUser)
	assert.Error(t, err)
}

func clearUsersDb() {
	usersDb.usersById = make(map[string]*User)
	usersDb.usersByEmail = make(map[string]*User)
}
