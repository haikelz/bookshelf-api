package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.CSRF())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

	e.Group("/api")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/v1", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World v1!")
	})

	e.GET("/v1/books", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World v2!")
	})

	e.GET("/v1/:bookId", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World v3!")
	})

	e.GET("/v1/books/:bookId", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World v4!")
	})

	e.POST("/v1/books", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World v5!")
	})

	e.PUT("/v1/:bookId", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World v5!")
	})

	e.DELETE("/v1/:bookId", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World v6!")
	})

	if err := e.Start(":5000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
