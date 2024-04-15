package session

import (
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

func TestCreateToken(t *testing.T) {
	tokenString, err := CreateSignedToken("0")
	if assert.NoError(t, err) {
		claims, err := GetClaims(tokenString)
		if assert.NoError(t, err) {
			assert.Equal(t, "0", claims.UserId)
		}
	}
}
