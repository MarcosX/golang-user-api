package api

import (
	"net/http"

	"github.com/brizenox/golang-user-api/internal/db"
	"github.com/brizenox/golang-user-api/internal/domain"
	"github.com/brizenox/golang-user-api/internal/session"
	"github.com/labstack/echo/v4"
)

type (
	userHandler struct {
		userRepository domain.UserRepository
	}
	putUserResponse struct {
		Token string
		User  *db.User
	}
)

func NewUserHanlder(userRepository domain.UserRepository) *userHandler {
	return &userHandler{
		userRepository: userRepository,
	}
}

func (handler *userHandler) getUser(c echo.Context) error {
	sessionClaims, err := session.ClaimsFromContext(c)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	user, err := handler.userRepository.GetUserByEmail(sessionClaims.Subject)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, user)
}

func (handler *userHandler) putUser(c echo.Context) error {
	sessionClaims, err := session.ClaimsFromContext(c)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	user, err := handler.userRepository.GetUserByEmail(sessionClaims.Subject)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	user, err = handler.userRepository.UpdateUser(user.Id, c.FormValue("name"), c.FormValue("email"), c.FormValue("password"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	token, err := session.SessionData().CreateSignedToken(user.Email)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, &putUserResponse{Token: token, User: user})
}
