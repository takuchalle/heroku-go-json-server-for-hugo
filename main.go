package main

import (
    "net/http"
    "os"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func main() {
    // Echo instance
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.Start(os.Getenv("PORT"))
}
