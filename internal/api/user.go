package api

import (
	"net/http"
	"strings"

	"github.com/brizenox/golang-user-api/internal/domain"
	"github.com/brizenox/golang-user-api/internal/session"
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
	sessionClaims, err := session.ClaimsFromContext(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid user session"})
	}

	user, err := u.userRepository.GetUser(userId)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if !strings.EqualFold(user.Email, sessionClaims.Subject) {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "invalid user session"})
	}

	return c.JSON(http.StatusOK, user)
}
