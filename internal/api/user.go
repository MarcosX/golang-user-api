package api

import (
	"log"
	"net/http"

	"github.com/brizenox/golang-user-api/internal/db"
	"github.com/labstack/echo/v4"
)

type (
	userHandler struct {
		userRepository    db.UserRepository
		sessionRepository db.SessionRepository
	}
)

func NewUserHanlder() *userHandler {
	return &userHandler{
		userRepository:    db.NewUserRepository(),
		sessionRepository: db.NewSessionRepository(),
	}
}

func (u *userHandler) getUser(c echo.Context) error {
	sessionId := c.Request().Header.Get("X-Session-ID")
	session, err := u.sessionRepository.GetSession(sessionId)
	if err != nil {
		return c.String(http.StatusBadRequest, "no user session")
	}
	userId := c.Param("id")
	sessionClaims, err := session.GetClaims()
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
