package api

import (
	"net/http"
	"strings"

	"github.com/brizenox/golang-user-api/internal/domain"
	"github.com/brizenox/golang-user-api/internal/session"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type (
	userHandler struct {
		userRepository domain.UserRepository
	}
)

func NewUserHanlder() *userHandler {
	return &userHandler{
		userRepository: domain.NewUserRepository(),
	}
}

func (u *userHandler) getUser(c echo.Context) error {
	userId := c.Param("id")
	if c.Get("user") == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid user session"})
	}
	sessionClaims := c.Get("user").(*jwt.Token).Claims.(*session.CustomClaims)

	user, err := u.userRepository.GetUser(userId)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if !strings.EqualFold(user.Email, sessionClaims.UserEmail) {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "invalid user session"})
	}

	return c.JSON(http.StatusOK, user)
}

func (u *userHandler) postUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := u.userRepository.CreateUser(name, email, password)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.JSON(http.StatusCreated, user)
}
