package api

import (
	"log"
	"net/http"

	"github.com/brizenox/golang-user-api/internal/db"
	"github.com/labstack/echo/v4"
)

type (
	userHandler struct {
		userDB db.UserRepository
	}
)

func NewUserHanlder() *userHandler {
	return &userHandler{
		userDB: db.NewUserRepository(),
	}
}

func (u *userHandler) getUser(c echo.Context) error {
	user, err := u.userDB.GetUser(c.Param("id"))
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, user)
}
