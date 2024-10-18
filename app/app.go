package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	mw "lol-team-maker/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(mw.LoggerConfig))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
