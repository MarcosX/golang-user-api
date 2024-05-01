package session

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
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

func TestClaimsFromContext(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/user/", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, &CustomClaims{})
	c.Set("Authorization", jwtToken)

	assert.NotNil(t, ClaimsFromContext(c))
}

func TestReadClaims(t *testing.T) {
	signedToken, _ := SessionData().CreateSignedToken("user@email.com")
	jwtToken, err := SessionData().ReadToken(signedToken)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, jwtToken)
	}
}

func TestSessionDataWithWrongPubKey(t *testing.T) {
	session = nil
	os.Setenv("SESSION_PUBLIC_KEY", "not/a/pubkey")

	assert.Panics(t, func() {
		SessionData()
	})

	os.Unsetenv("SESSION_PUBLIC_KEY")
}

func TestSessionDataWithWrongPrivKey(t *testing.T) {
	session = nil
	os.Setenv("SESSION_PRIVATE_KEY", "not/a/privkey")

	assert.Panics(t, func() {
		SessionData()
	})

	os.Unsetenv("SESSION_PRIVATE_KEY")
}
