package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func (app *Application) registerHeartbeat(c echo.Context) error {
	device := c.Request().Header.Get("Device")
	if device == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "You suck!")
	}

	app.DoSomething(c.Request().Context())

	return c.String(http.StatusOK, fmt.Sprintf("Hello %s", device))
}
