package routes

import (
	"bookshelf/internal/controllers"
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

func BookRoute(db *gorm.DB, e *echo.Echo) *echo.Echo {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Prometheus
	e.Group("/api").GET("/metrics", echoprometheus.NewHandler())

	e.Group("/api").GET("/v1/books", func(c echo.Context) error {
		return controllers.BooksController(c, db)
	})

	e.Group("/api").GET("/v1/books/:bookId", func(c echo.Context) error {
		return controllers.BookByIdController(c, db)
	})

	e.Group("/api").POST("/v1/books", func(c echo.Context) error {
		return controllers.CreateBookController(c, db)
	})

	e.Group("/api").PUT("/v1/books/:bookId", func(c echo.Context) error {
		return controllers.UpdateBookController(c, db)
	})

	e.Group("/api").DELETE("/v1/books/:bookId", func(c echo.Context) error {
		return controllers.DeleteBookController(c, db)
	})

	// Swagger
	e.Group("/api").GET("/v1/docs/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))

	return e
}
