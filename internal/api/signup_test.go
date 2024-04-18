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

func TestPostSignup(t *testing.T) {
	form := url.Values{}
	form.Add("email", "usertest@email.com")
	form.Add("password", "pass")
	form.Add("name", "User")
	req := httptest.NewRequest(http.MethodPost, "/signup", nil)
	req.Form = form
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	echoContext := echo.New().NewContext(req, rec)

	h := NewSignupHandlerTest()
	if assert.NoError(t, h.postSignup(echoContext)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		response := &signupResponse{}
		json.NewDecoder(rec.Body).Decode(response)
		sessionData, err := session.SessionData().ReadClaims(response.Token)
		if assert.NoError(t, err) {
			assert.Equal(t, "usertest@email.com", sessionData.Subject)
		}
	}
}

func TestPostSignupFailedToCreateUser(t *testing.T) {
	form := url.Values{}
	form.Add("email", "usertest@email.com")
	form.Add("password", "pass")
	form.Add("name", "Crash()")
	req := httptest.NewRequest(http.MethodPost, "/signup", nil)
	req.Form = form
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	echoContext := echo.New().NewContext(req, rec)

	h := NewSignupHandlerTest()
	if assert.NoError(t, h.postSignup(echoContext)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		_, err := h.userRepository.GetUserByEmail("usertest@email.com")
		assert.Errorf(t, err, "usertest@email.com")
	}
}

func NewSignupHandlerTest() *signupHandler {
	return &signupHandler{
		userRepository: domain.NewMockUserRepository(),
	}
}
