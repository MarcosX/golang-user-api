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

func NewLoginHandler(userRepository domain.UserRepository) *loginHandler {
	return &loginHandler{
		userRepository: userRepository,
	}
}

func (h *loginHandler) postLogin(c echo.Context) error {
	emailFromForm := c.FormValue("email")
	passwordFromForm := c.FormValue("password")
	if emailFromForm == "" || passwordFromForm == "" {
		log.Println("email or password is empty")
		return c.NoContent(http.StatusBadRequest)
	}
	user, err := h.userRepository.GetUserByEmail(emailFromForm)
	if err != nil {
		log.Println(err)
		return c.NoContent(http.StatusUnauthorized)
	}
	if user.PasswordMatches(passwordFromForm) {
		token, err := session.SessionData().CreateSignedToken("user@email.com")
		if err != nil {
			log.Println(err)
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, &loginResponse{Token: token})
	}
	return c.NoContent(http.StatusUnauthorized)
}
