package routes

import (
	"bookshelf/internal/configs"
	"bookshelf/internal/controllers"
	"bookshelf/internal/entities"
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func BookRoute() *echo.Echo {
	app := echo.New()
	db := configs.NewGorm()

	err := db.AutoMigrate(&entities.Book{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	app.Group("/api").GET("/metrics", echoprometheus.NewHandler())

	app.Group("/api").GET("/v1/books", func(c echo.Context) error {
		return controllers.BooksController(c, db)
	})

	app.Group("/api").GET("/v1/books/:bookId", func(c echo.Context) error {
		return controllers.BookByIdController(c, db)
	})

	app.Group("/api").POST("/v1/books", func(c echo.Context) error {
		return controllers.CreateBookController(c, db)
	})

	app.Group("/api").PUT("/v1/books/:bookId", func(c echo.Context) error {
		return controllers.UpdateBookController(c, db)
	})

	app.Group("/api").DELETE("/v1/books/:bookId", func(c echo.Context) error {
		return controllers.DeleteBookController(c, db)
	})

	app.Group("/api").GET("/v1/docs/*", echoSwagger.WrapHandler)

	app.Logger.Fatal(app.Start(":8080"))

	return app
}
