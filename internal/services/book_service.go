package services

import (
	"bookshelf/internal/entities"
	"bookshelf/internal/repositories"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetBooks(c context.Context, books []entities.Book, db *gorm.DB) (entities.BooksResponse, error) {
	err := repositories.GetAll(db, &books)
	if err != nil {
		return entities.BooksResponse{}, err
	}

	return entities.BooksResponse{
		Status: "success",
		Data:   books,
	}, nil
}

func GetBookById(c context.Context, id string, book entities.Book, db *gorm.DB) (entities.BookDetailResponse, error) {
	err := repositories.GetById(db, &book, id)
	if err != nil {
		return entities.BookDetailResponse{}, err
	}

	return entities.BookDetailResponse{
		Status: "success",
		Data:   book,
	}, nil
}

func CreateBook(c context.Context, body io.ReadCloser, book entities.Book, db *gorm.DB) (entities.BookChangesResponse, error) {
	json.NewDecoder(body).Decode(&book)
	book.ID = uuid.New().String()

	err := repositories.Create(db, &book)
	if err != nil {
		return entities.BookChangesResponse{}, err
	}

	return entities.BookChangesResponse{
		Status:  "success",
		Message: "Book created successfully",
	}, nil
}

func UpdateBook(c context.Context, id string, body io.ReadCloser, book entities.Book, db *gorm.DB) (entities.BookChangesResponse, error) {
	fmt.Println(id)
	json.NewDecoder(body).Decode(&book)
	fmt.Println(book)

	err := repositories.Update(db, &book, id)
	if err != nil {
		return entities.BookChangesResponse{}, err
	}

	return entities.BookChangesResponse{
		Status:  "success",
		Message: "Book updated successfully",
	}, nil
}

func DeleteBook(c context.Context, id string, book entities.Book, db *gorm.DB) (entities.BookChangesResponse, error) {
	err := repositories.Delete(db, &book, id)
	if err != nil {
		return entities.BookChangesResponse{}, err
	}

	return entities.BookChangesResponse{
		Status:  "success",
		Message: "Book deleted successfully",
	}, nil
}
