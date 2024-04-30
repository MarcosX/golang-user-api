package api

import (
	"net/http"

	"github.com/brizenox/golang-user-api/internal/domain"
	"github.com/brizenox/golang-user-api/internal/session"
	"github.com/labstack/echo/v4"
)

type (
	loginHandler struct {
		userRepository domain.UserRepository
	}
	loginResponse struct {
		Token string
	}
)

func NewLoginHandler(userRepository domain.UserRepository) *loginHandler {
	return &loginHandler{
		userRepository: userRepository,
	}
}

func (handler *loginHandler) postLogin(c echo.Context) error {
	user, err := handler.userRepository.GetUserByEmail(c.FormValue("email"))
	if err != nil || !user.PasswordMatches(c.FormValue("password")) {
		return c.NoContent(http.StatusUnauthorized)
	}

	token, err := session.SessionData().CreateSignedToken(user.Email)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, &loginResponse{Token: token})
}
