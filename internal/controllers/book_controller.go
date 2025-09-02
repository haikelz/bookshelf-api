package controllers

import (
	"bookshelf/internal/entities"
	"bookshelf/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BooksController(c echo.Context, db *gorm.DB) error {
	books := []entities.Book{}
	response, err := services.GetBooks(c.Request().Context(), books, db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func BookByIdController(c echo.Context, db *gorm.DB) error {
	book := entities.Book{}
	response, err := services.GetBookById(c.Request().Context(), c.Param("bookId"), book, db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, response)
}

func CreateBookController(c echo.Context, db *gorm.DB) error {
	book := entities.Book{}
	response, err := services.CreateBook(c.Request().Context(), c.Request().Body, book, db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, response)
}

func UpdateBookController(c echo.Context, db *gorm.DB) error {
	book := entities.Book{}
	response, err := services.UpdateBook(c.Request().Context(), c.Param("id"), c.Request().Body, book, db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteBookController(c echo.Context, db *gorm.DB) error {
	book := entities.Book{}
	response, err := services.DeleteBook(c.Request().Context(), c.Param("bookId"), book, db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, response)
}
