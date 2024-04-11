package session

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetclaims(t *testing.T) {
	claims, err := GetClaims(ValidSession)
	if assert.NoError(t, err) {
		assert.Equal(t, "0", claims.UserId)
	}
}

func TestWrongSecretSessionToken(t *testing.T) {
	claims, err := GetClaims(WrongSecretSession)
	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestGetTestSessionSecretKey(t *testing.T) {
	key := getSessionKey()
	assert.Equal(t, []byte("Session Secret Key"), key)
}

func TestGetSessionSecretKeyFromEnv(t *testing.T) {
	os.Setenv("SESSION_KEY", "Key From Env")
	key := getSessionKey()
	assert.Equal(t, []byte("Key From Env"), key)
	os.Unsetenv("SESSION_KEY")
}
