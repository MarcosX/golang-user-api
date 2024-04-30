package api

import (
	"net/http"

	"github.com/brizenox/golang-user-api/internal/db"
	"github.com/brizenox/golang-user-api/internal/domain"
	"github.com/brizenox/golang-user-api/internal/session"
	"github.com/labstack/echo/v4"
)

type (
	signupHandler struct {
		userRepository domain.UserRepository
	}
	signupResponse struct {
		Token string
		User  *db.User
	}
)

func NewSignupHandler(userRepository domain.UserRepository) *signupHandler {
	return &signupHandler{
		userRepository: userRepository,
	}
}

func (handler *signupHandler) postSignup(c echo.Context) error {
	user, err := handler.userRepository.CreateUser(c.FormValue("name"), c.FormValue("email"), c.FormValue("password"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	token, err := session.SessionData().CreateSignedToken(user.Email)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, &signupResponse{Token: token, User: user})
}
