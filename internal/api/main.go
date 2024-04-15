package api

import (
	"github.com/brizenox/golang-user-api/internal/session"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupHandlers(echoInstance *echo.Echo) {
	echoInstance.Pre(middleware.RemoveTrailingSlash())
	echoInstance.Use(middleware.Recover())

	echoInstance.GET("/health", getHealth)

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(session.CustomClaims)
		},
		SigningKey: session.GetPublicSessionKey(),
	}
	userGroup := echoInstance.Group("/user", echojwt.WithConfig(config))

	userHandler := NewUserHanlder()
	userGroup.GET("/:id", userHandler.getUser)
}
