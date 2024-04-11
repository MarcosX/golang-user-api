package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSession(t *testing.T) {
	sessionRepository := NewMockSessionRepository()
	session, err := sessionRepository.GetSession("samplesession")
	if assert.NoError(t, err) {
		assert.NotNil(t, session)
		claims, err := session.GetClaims()
		if assert.NoError(t, err) {
			assert.Equal(t, "0", claims.UserId)
		}
	}
}

func TestWrongSecretSessionToken(t *testing.T) {
	sessionRepository := NewMockSessionRepository()
	session, err := sessionRepository.GetSession("wrongsecretsession")
	if assert.NoError(t, err) {
		_, err := session.GetClaims()
		assert.Error(t, err)
	}
}
