package api

import (
	"github.com/brizenox/golang-user-api/internal/api/health"
	"github.com/labstack/echo/v4"
)

func SetupHandlers(echoInstance *echo.Echo) {
	echoInstance.GET("/health", health.Health)
}
