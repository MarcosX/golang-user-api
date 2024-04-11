package api

import (
	"log"
	"net/http"

	"github.com/brizenox/golang-user-api/internal/db"
	"github.com/brizenox/golang-user-api/internal/session"
	"github.com/labstack/echo/v4"
)

type (
	userHandler struct {
		userRepository db.UserRepository
	}
)

func NewUserHanlder() *userHandler {
	return &userHandler{
		userRepository: db.NewUserRepository(),
	}
}

func (u *userHandler) getUser(c echo.Context) error {
	sessionToken := c.Request().Header.Get("X-Session-ID")
	if sessionToken == "" {
		return c.String(http.StatusBadRequest, "no user session")
	}

	userId := c.Param("id")
	sessionClaims, err := session.GetClaims(sessionToken)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if sessionClaims.UserId == userId {
		user, err := u.userRepository.GetUser(userId)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusNotFound, nil)
		}
		return c.JSON(http.StatusOK, user)
	}
	return c.String(http.StatusBadRequest, "invalid user session")
}
