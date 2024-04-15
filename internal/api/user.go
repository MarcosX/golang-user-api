package api

import (
	"log"
	"net/http"

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

	if sessionClaims.UserId == userId {
		user, err := u.userRepository.GetUser(userId)
		if err != nil {
			log.Println(err)
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusForbidden, map[string]string{"message": "invalid user session"})
}
