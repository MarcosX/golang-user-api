package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/brizenox/golang-user-api/internal/domain"
	"github.com/brizenox/golang-user-api/internal/session"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	b, _ := json.Marshal(domain.UserSample)
	expectedResponse := string(b)

	req := httptest.NewRequest(http.MethodGet, "/user/", nil)
	rec := httptest.NewRecorder()

	echoContext := echo.New().NewContext(req, rec)
	signedToken, _ := session.SessionData().CreateSignedToken("user@email.com")
	jwtToken, _ := session.SessionData().ReadToken(signedToken)
	echoContext.Set("Authorization", jwtToken)

	handler := newUserHandlerTest()

	if assert.NoError(t, handler.getUser(echoContext)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedResponse, strings.Trim(rec.Body.String(), "\n"))
	}
}

func TestGetUserNonExistingSession(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/user/", nil)
	rec := httptest.NewRecorder()

	echoContext := echo.New().NewContext(req, rec)

	handler := newUserHandlerTest()

	if assert.NoError(t, handler.getUser(echoContext)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func newUserHandlerTest() *userHandler {
	return &userHandler{
		userRepository: domain.NewMockUserRepository(),
	}
}
