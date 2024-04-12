package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/brizenox/golang-user-api/internal/db"
	"github.com/brizenox/golang-user-api/internal/session"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	b, _ := json.Marshal(db.UserSample)
	expectedResponse := string(b)

	req := httptest.NewRequest(http.MethodGet, "/user/:id", nil)
	rec := httptest.NewRecorder()

	echoContext := echo.New().NewContext(req, rec)
	echoContext.Set("user", createToken(session.ValidSession))
	echoContext.SetParamNames("id")
	echoContext.SetParamValues("0")

	handler := newUserHandlerTest()

	if assert.NoError(t, handler.getUser(echoContext)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedResponse, strings.Trim(rec.Body.String(), "\n"))
	}
}

func TestGetNonExistingUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/user/:id", nil)
	rec := httptest.NewRecorder()

	echoContext := echo.New().NewContext(req, rec)
	echoContext.Set("user", createToken(session.InvalidUserValidSession))
	echoContext.SetParamNames("id")
	echoContext.SetParamValues("99")

	handler := newUserHandlerTest()

	if assert.NoError(t, handler.getUser(echoContext)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestGetUserNonExistingSession(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/user/:id", nil)
	rec := httptest.NewRecorder()

	echoContext := echo.New().NewContext(req, rec)
	echoContext.SetParamNames("id")
	echoContext.SetParamValues("0")

	handler := newUserHandlerTest()

	if assert.NoError(t, handler.getUser(echoContext)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestGetUserInvalidSessionUserId(t *testing.T) {
	echo := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user/:id", nil)
	rec := httptest.NewRecorder()

	echoContext := echo.NewContext(req, rec)
	echoContext.Set("user", createToken(session.ValidSession))
	echoContext.SetParamNames("id")
	echoContext.SetParamValues("99")

	handler := newUserHandlerTest()

	if assert.NoError(t, handler.getUser(echoContext)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func newUserHandlerTest() *userHandler {
	return &userHandler{
		userRepository: db.NewMockUserRepository(),
	}
}

func createToken(signedToken string) *jwt.Token {
	jwtToken, _ := jwt.ParseWithClaims(signedToken, &session.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return session.GetSessionKey(), nil
	})
	return jwtToken
}
