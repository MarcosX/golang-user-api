package api

import (
	"github.com/brizenox/golang-user-api/internal/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupHandlers(echoInstance *echo.Echo) {
	echoInstance.Pre(middleware.RemoveTrailingSlash())
	echoInstance.Use(middleware.Recover())

	echoInstance.GET("/health", getHealth)
	userGroup := echoInstance.Group("/user", session.EnforceValidSession())

	userHandler := NewUserHanlder()
	userGroup.GET("/:id", userHandler.getUser)

	loginHandler := NewLoginHandler()
	loginGroup := echoInstance.Group("/login")
	loginGroup.POST("", loginHandler.postLogin)

	signupHandler := NewSignupHandler()
	signupGroup := echoInstance.Group("/signup")
	signupGroup.POST("", signupHandler.postSignup)
}
