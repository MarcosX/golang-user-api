package api

import (
	"github.com/labstack/echo/v4"
)

func SetupHandlers(echoInstance *echo.Echo) {
	echoInstance.GET("/health", getHealth)

	userHandler := NewUserHanlder()
	echoInstance.GET("/user/:id", userHandler.getUser)
}
