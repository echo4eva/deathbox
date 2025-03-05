package main

import "github.com/labstack/echo"

func (app *Application) InitRouting(e *echo.Echo) {
	app.e.POST("/heartbeat", app.registerHeartbeat)
}
