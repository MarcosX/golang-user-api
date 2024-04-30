package api

import (
	"github.com/brizenox/golang-user-api/internal/domain"
	"github.com/brizenox/golang-user-api/internal/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupHandlers(echoInstance *echo.Echo) {
	userRepository := domain.NewUserRepository()

	echoInstance.Pre(middleware.RemoveTrailingSlash())
	echoInstance.Use(middleware.Recover())

	echoInstance.GET("/health", getHealth)

	userGroup := echoInstance.Group("/user", session.EnforceValidSession())
	userHandler := NewUserHanlder(userRepository)
	userGroup.GET("", userHandler.getUser)
	userGroup.PUT("", userHandler.putUser)

	loginHandler := NewLoginHandler(userRepository)
	loginGroup := echoInstance.Group("/login")
	loginGroup.POST("", loginHandler.postLogin)

	signupHandler := NewSignupHandler(userRepository)
	signupGroup := echoInstance.Group("/signup")
	signupGroup.POST("", signupHandler.postSignup)
}
