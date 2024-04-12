package main

import (
	"fmt"

	"github.com/brizenox/golang-user-api/internal/api"
	"github.com/labstack/echo/v4"
)

const port = 8080

func main() {
	echoInstance := echo.New()
	api.SetupHandlers(echoInstance)

	echoInstance.Logger.Fatal(echoInstance.Start(fmt.Sprintf(":%d", port)))
}
