package controllers

import (
	"bookshelf/internal/entities"
	"bookshelf/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// BooksController godoc
// @Summary Get all books
// @Description Get all books
// @Accept json
// @Produce json
// @Success 200 {object} entities.BooksResponse
// @Failure 500 {object} entities.BookChangesResponse
// @Router /api/v1/books [get]
// @Tags books
func BooksController(c echo.Context, db *gorm.DB) error {
	var books []entities.Book

	response, err := services.GetBooks(c.Request().Context(), books, db)
	if err != nil {
		response := entities.BookChangesResponse{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, response)
}

// BookByIdController godoc
// @Summary Get book by id
// @Description Get book by id
// @Accept json
// @Produce json
// @Success 200 {object} entities.BookDetailResponse
// @param bookId path string true "Book ID"
// @Failure 500 {object} entities.BookChangesResponse
// @Router /api/v1/books/{bookId} [get]
// @Tags books
func BookByIdController(c echo.Context, db *gorm.DB) error {
	var book entities.Book

	response, err := services.GetBookById(c.Request().Context(), c.Param("bookId"), book, db)
	if err != nil {
		response := entities.BookChangesResponse{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, response)
}

// CreateBookController godoc
// @Summary Create book
// @Description Create book
// @Accept json
// @Produce json
// @Success 200 {object} entities.BookChangesResponse
// @Param book body entities.Book true "Book"
// @Failure 500 {object} entities.BookChangesResponse
// @Router /api/v1/books [post]
// @Tags books
func CreateBookController(c echo.Context, db *gorm.DB) error {
	var book entities.Book

	response, err := services.CreateBook(c.Request().Context(), c.Request().Body, book, db)
	if err != nil {
		response := entities.BookChangesResponse{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, response)
}

// UpdateBookController godoc
// @Summary Update book
// @Description Update book
// @Accept json
// @Produce json
// @Success 200 {object} entities.BookChangesResponse
// @Failure 500 {object} entities.BookChangesResponse
// @Router /api/v1/books/{bookId} [put]
// @Tags books
func UpdateBookController(c echo.Context, db *gorm.DB) error {
	var book entities.Book

	response, err := services.UpdateBook(c.Request().Context(), c.Param("bookId"), c.Request().Body, book, db)
	if err != nil {
		response := entities.BookChangesResponse{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, response)
}

// DeleteBookController godoc
// @Summary Delete book
// @Description Delete book
// @Accept json
// @Produce json
// @Success 200 {object} entities.BookChangesResponse
// @Param bookId path string true "Book ID"
// @Failure 500 {object} entities.BookChangesResponse
// @Router /api/v1/books/{bookId} [delete]
// @Tags books
func DeleteBookController(c echo.Context, db *gorm.DB) error {
	var book entities.Book

	response, err := services.DeleteBook(c.Request().Context(), c.Param("bookId"), book, db)
	if err != nil {
		response := entities.BookChangesResponse{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, response)
}
