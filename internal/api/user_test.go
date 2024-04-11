package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/brizenox/golang-user-api/internal/db"
	"github.com/brizenox/golang-user-api/internal/session"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	b, _ := json.Marshal(db.UserSample)
	expectedResponse := string(b)

	echo := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user/:id", nil)
	req.Header.Set("X-Session-ID", session.ValidSession)
	rec := httptest.NewRecorder()
	echoContext := echo.NewContext(req, rec)
	echoContext.SetParamNames("id")
	echoContext.SetParamValues("0")

	handler := newUserHandlerTest()

	if assert.NoError(t, handler.getUser(echoContext)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedResponse, strings.Trim(rec.Body.String(), "\n"))
	}
}

func TestGetNonExistingUser(t *testing.T) {
	echo := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user/:id", nil)
	req.Header.Set("X-Session-ID", session.InvalidUserValidSession)
	rec := httptest.NewRecorder()
	echoContext := echo.NewContext(req, rec)
	echoContext.SetParamNames("id")
	echoContext.SetParamValues("99")

	handler := newUserHandlerTest()

	if assert.NoError(t, handler.getUser(echoContext)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestGetUserNonExistingSession(t *testing.T) {
	echo := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user/:id", nil)
	rec := httptest.NewRecorder()
	echoContext := echo.NewContext(req, rec)
	echoContext.SetParamNames("id")
	echoContext.SetParamValues("0")

	handler := newUserHandlerTest()

	if assert.NoError(t, handler.getUser(echoContext)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestGetUserInvalidSession(t *testing.T) {
	echo := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user/:id", nil)
	req.Header.Set("X-Session-ID", session.WrongSecretSession)
	rec := httptest.NewRecorder()
	echoContext := echo.NewContext(req, rec)
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
	req.Header.Set("X-Session-ID", session.ValidSession)
	rec := httptest.NewRecorder()
	echoContext := echo.NewContext(req, rec)
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
