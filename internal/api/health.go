package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthResponse struct {
	Status string `json:"status"`
}

func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, &healthResponse{Status: "UP"})
}
