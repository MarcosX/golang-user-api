package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/brizenox/golang-user-api/internal/domain"
	"github.com/brizenox/golang-user-api/internal/session"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	b, _ := json.Marshal(domain.UserSample)
	expectedResponse := string(b)

	req := httptest.NewRequest(http.MethodGet, "/user/:id", nil)
	rec := httptest.NewRecorder()

	echoContext := echo.New().NewContext(req, rec)
	echoContext.Set("user", jwt.NewWithClaims(jwt.SigningMethodRS256, &session.CustomClaims{UserId: "0"}))
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
	echoContext.Set("user", jwt.NewWithClaims(jwt.SigningMethodRS256, &session.CustomClaims{UserId: "99"}))
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
	req := httptest.NewRequest(http.MethodGet, "/user/:id", nil)
	rec := httptest.NewRecorder()

	echoContext := echo.New().NewContext(req, rec)
	echoContext.Set("user", jwt.NewWithClaims(jwt.SigningMethodRS256, &session.CustomClaims{UserId: "0"}))
	echoContext.SetParamNames("id")
	echoContext.SetParamValues("99")

	handler := newUserHandlerTest()

	if assert.NoError(t, handler.getUser(echoContext)) {
		assert.Equal(t, http.StatusForbidden, rec.Code)
	}
}

func newUserHandlerTest() *userHandler {
	return &userHandler{
		userRepository: domain.NewMockUserRepository(),
	}
}
