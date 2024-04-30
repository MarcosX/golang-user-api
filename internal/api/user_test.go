package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
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

func TestGetUserNonExistingUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/user/", nil)
	rec := httptest.NewRecorder()

	echoContext := echo.New().NewContext(req, rec)
	signedToken, _ := session.SessionData().CreateSignedToken("not.a.user@email.com")
	jwtToken, _ := session.SessionData().ReadToken(signedToken)
	echoContext.Set("Authorization", jwtToken)

	handler := newUserHandlerTest()

	if assert.NoError(t, handler.getUser(echoContext)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestPutUser(t *testing.T) {
	form := url.Values{}
	form.Add("email", "updated@email.com")
	form.Add("password", "updatedpass")
	form.Add("name", "Updated Name")
	req := httptest.NewRequest(http.MethodPut, "/user/", nil)
	req.Form = form
	rec := httptest.NewRecorder()

	handler := newUserHandlerTest()
	user, _ := handler.userRepository.CreateUser("User", "user@email.com", "pass")
	signedToken, _ := session.SessionData().CreateSignedToken(user.Email)
	jwtToken, _ := session.SessionData().ReadToken(signedToken)

	echoContext := echo.New().NewContext(req, rec)
	echoContext.Set("Authorization", jwtToken)

	if assert.NoError(t, handler.putUser(echoContext)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		response := &putUserResponse{}
		json.NewDecoder(rec.Body).Decode(response)
		token, err := session.SessionData().ReadClaims(response.Token)
		if assert.NoError(t, err) {
			assert.Equal(t, "updated@email.com", token.Subject)
		}
	}
}

func TestPutNonExistingUser(t *testing.T) {
	form := url.Values{}
	form.Add("email", "updated@email.com")
	form.Add("password", "updatedpass")
	form.Add("name", "Updated Name")
	req := httptest.NewRequest(http.MethodPut, "/user/", nil)
	req.Form = form
	rec := httptest.NewRecorder()

	handler := newUserHandlerTest()
	signedToken, _ := session.SessionData().CreateSignedToken("not.a.user@email.com")
	jwtToken, _ := session.SessionData().ReadToken(signedToken)

	echoContext := echo.New().NewContext(req, rec)
	echoContext.Set("Authorization", jwtToken)

	if assert.NoError(t, handler.putUser(echoContext)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestPutFailToSaveUser(t *testing.T) {
	form := url.Values{}
	form.Add("email", "updated@email.com")
	form.Add("password", "updatedpass")
	form.Add("name", "Crash()")
	req := httptest.NewRequest(http.MethodPut, "/user/", nil)
	req.Form = form
	rec := httptest.NewRecorder()

	handler := newUserHandlerTest()
	user, _ := handler.userRepository.CreateUser("User", "user@email.com", "pass")
	signedToken, _ := session.SessionData().CreateSignedToken(user.Email)
	jwtToken, _ := session.SessionData().ReadToken(signedToken)

	echoContext := echo.New().NewContext(req, rec)
	echoContext.Set("Authorization", jwtToken)

	if assert.NoError(t, handler.putUser(echoContext)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func newUserHandlerTest() *userHandler {
	return &userHandler{
		userRepository: domain.NewMockUserRepository(),
	}
}
