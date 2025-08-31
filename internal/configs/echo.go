package configs

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewEcho() *echo.Echo {
	e := echo.New()

	e.Use(echoprometheus.NewMiddleware("bookshelf"))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.CSRF())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())

	return e
}
