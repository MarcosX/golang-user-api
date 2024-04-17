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

func NewSignupHandler() *signupHandler {
	return &signupHandler{
		userRepository: domain.NewUserRepository(),
	}
}

func (h *signupHandler) postSignup(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := h.userRepository.CreateUser(name, email, password)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	token, err := session.SessionData().CreateSignedToken(email)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, &signupResponse{Token: token, User: user})
}
