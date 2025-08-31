package main

import (
	"bookshelf/internal/configs"
	"bookshelf/internal/controllers"
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

func main() {
	app := configs.NewEcho()

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	app.Group("/api").GET("/metrics", echoprometheus.NewHandler())

	app.Group("/api").GET("/v1/books", controllers.BooksController)

	app.Group("/api").GET("/v1/books/:bookId", controllers.BookByIdController)

	app.Group("/api").POST("/v1/books", controllers.CreateBookController)

	app.Group("/api").PUT("/v1/books/:bookId", controllers.UpdateBookController)

	app.Group("/api").DELETE("/v1/books/:bookId", controllers.DeleteBookController)

	app.Group("/api").GET("/v1/docs/*", echoSwagger.WrapHandler)

	if err := app.Start(":5000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
