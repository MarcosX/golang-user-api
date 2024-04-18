package session

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetclaims(t *testing.T) {
	tokenString, _ := SessionData().CreateSignedToken("user@email.com")
	claims, err := SessionData().ReadClaims(tokenString)
	if assert.NoError(t, err) {
		assert.Equal(t, "user@email.com", claims.Subject)
	}
}

func TestWrongSecretSessionToken(t *testing.T) {
	fakeSessionToke := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIwIn0.nL3zQibYzBCqwzILJ6KJQSiYEEXjxqnu5rM0_U-ZH0E"
	claims, err := SessionData().ReadClaims(fakeSessionToke)
	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestCreateToken(t *testing.T) {
	tokenString, err := SessionData().CreateSignedToken("user@email.com")
	if assert.NoError(t, err) {
		claims, err := SessionData().ReadClaims(tokenString)
		if assert.NoError(t, err) {
			assert.Equal(t, "user@email.com", claims.Subject)
		}
	}
}
