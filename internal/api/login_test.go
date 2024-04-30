package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/brizenox/golang-user-api/internal/domain"
	"github.com/brizenox/golang-user-api/internal/session"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPostLogin(t *testing.T) {
	form := url.Values{}
	form.Add("email", "user@email.com")
	form.Add("password", "pass")
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	req.Form = form
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	echoContext := echo.New().NewContext(req, rec)

	h := newLoginHandlerTest()
	h.userRepository.CreateUser("User", "user@email.com", "pass")
	if assert.NoError(t, h.postLogin(echoContext)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		response := &loginResponse{}
		json.NewDecoder(rec.Body).Decode(response)
		sessionData, err := session.SessionData().ReadClaims(response.Token)
		if assert.NoError(t, err) {
			assert.Equal(t, "user@email.com", sessionData.Subject)
		}
	}
}

func TestPostLoginWrongPassword(t *testing.T) {
	form := url.Values{}
	form.Add("email", "user@email.com")
	form.Add("password", "admin123")
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	req.Form = form
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	echoContext := echo.New().NewContext(req, rec)

	h := newLoginHandlerTest()
	if assert.NoError(t, h.postLogin(echoContext)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	}
}

func TestPostLoginInvalidUser(t *testing.T) {
	form := url.Values{}
	form.Add("email", "not.a.user@email.com")
	form.Add("password", "admin123")
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	req.Form = form
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	echoContext := echo.New().NewContext(req, rec)

	h := newLoginHandlerTest()
	if assert.NoError(t, h.postLogin(echoContext)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	}
}

func newLoginHandlerTest() *loginHandler {
	return &loginHandler{
		userRepository: domain.NewMockUserRepository(),
	}
}
