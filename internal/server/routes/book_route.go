package routes

import (
	"bookshelf-api/internal/controllers"
	"bookshelf-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

func BookRoute(api fiber.Router) {
	route := api.Group("/api")

	bookService := services.NewBookService()
	bookController := controllers.NewBookController()

	route.Get("/v1")
	route.Get("/v1/books")
	route.Get("/v1/:bookId")
	route.Post("/v1/books")
	route.Put("/v1/:bookId")
	route.Delete("/v1/:bookId")
}
