package main

import (
	"fmt"

	"github.com/brizenox/golang-user-api/internal/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const port = 8080

func main() {
	echoInstance := echo.New()
	echoInstance.Pre(middleware.RemoveTrailingSlash())
	echoInstance.Use(middleware.Recover())

	echoInstance.GET("/health", api.Health)

	echoInstance.Logger.Fatal(echoInstance.Start(fmt.Sprintf(":%d", port)))
}
