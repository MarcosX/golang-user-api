package api

import (
	"log"
	"net/http"

	"github.com/brizenox/golang-user-api/internal/domain"
	"github.com/brizenox/golang-user-api/internal/session"
	"github.com/labstack/echo/v4"
)

type (
	loginHandler struct {
		userRepository domain.UserRepository
	}
	loginResponse struct {
		Token string
	}
)

func NewLoginHandler() *loginHandler {
	return &loginHandler{
		userRepository: domain.NewUserRepository(),
	}
}

func (h *loginHandler) postLogin(c echo.Context) error {
	user, err := h.userRepository.GetUserByEmail(c.FormValue("email"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if user.PasswordMatches(c.FormValue("password")) {
		token, err := session.CreateSignedToken("0")
		if err != nil {
			log.Println(err)
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, &loginResponse{Token: token})
	}
	return c.NoContent(http.StatusUnauthorized)
}
