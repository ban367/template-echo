package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	initRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func initRoutes(e *echo.Echo) {
	e.GET("/", hello)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
