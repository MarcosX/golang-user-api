package api

import (
	"net/http"

	"github.com/brizenox/golang-user-api/internal/db"
	"github.com/labstack/echo/v4"
)

type (
	userResponse struct {
		Email       string `json:"email"`
		DisplayName string `json:"display_name"`
	}

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
	userDb, err := u.userDB.GetUser(c.Param("id"))
	if err != nil {
		return err
	}
	user := &userResponse{
		Email:       userDb.Email,
		DisplayName: userDb.Name,
	}
	return c.JSON(http.StatusOK, user)
}
