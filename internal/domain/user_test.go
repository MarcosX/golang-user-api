package domain

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
