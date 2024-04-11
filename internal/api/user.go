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
	_, err := u.sessionRepository.GetSession(sessionId)
	if err != nil {
		return err
	}
	// session.Token.Claims.

	user, err := u.userRepository.GetUser(c.Param("id"))
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, user)
}
