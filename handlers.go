package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func handleHeartbeat(c echo.Context) error {
	device := c.Request().Header.Get("Device")
	clientSecret := c.Request().Header.Get("Secret")
	if device == "" || clientSecret != os.Getenv("SECRET") {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	registerHeartbeat(c.Request().Context(), device)

	return c.String(http.StatusOK, fmt.Sprintf("Hello %s", device))
}
