package controllers

import (
	"bookshelf/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func BooksController(c echo.Context) error {
	books := services.GetBooks(c.Request().Context())
	return c.JSON(http.StatusOK, books)
}

func BookByIdController(c echo.Context) error {
	book := services.GetBookById(c.Request().Context(), c.Param("id"))
	return c.JSON(http.StatusOK, book)
}

func CreateBookController(c echo.Context) error {
	book := services.CreateBook(c.Request().Context(), c.Request().Body)
	return c.JSON(http.StatusOK, book)
}

func UpdateBookController(c echo.Context) error {
	book := services.UpdateBook(c.Request().Context(), c.Param("id"), c.Request().Body)
	return c.JSON(http.StatusOK, book)
}

func DeleteBookController(c echo.Context) error {
	book := services.DeleteBook(c.Request().Context(), c.Param("id"))
	return c.JSON(http.StatusOK, book)
}
