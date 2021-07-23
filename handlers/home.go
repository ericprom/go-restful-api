package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomePage(c echo.Context) error {
	return c.String(http.StatusOK, "Homepage Endpoint Hit")
}
