package services

import (
	"bookshelf/internal/entities"
	"bookshelf/internal/repositories"
	"context"
	"encoding/json"
	"io"
	"slices"

	"github.com/google/uuid"
)

func GetBooks(c context.Context) entities.BooksResponse {
	return entities.BooksResponse{
		Status: "success",
		Data:   repositories.Books,
	}
}

func GetBookById(c context.Context, id string) entities.BookDetailResponse {
	var book entities.Book

	for _, b := range repositories.Books {
		if b.ID == id {
			book = b
			break
		}
	}

	return entities.BookDetailResponse{
		Status: "success",
		Data:   book,
	}
}

func CreateBook(c context.Context, body io.ReadCloser) entities.BookChangesResponse {
	book := entities.Book{}
	json.NewDecoder(body).Decode(&book)

	book.ID = uuid.New().String()

	repositories.Books = append(repositories.Books, book)

	return entities.BookChangesResponse{
		Status:  "success",
		Message: "Book created successfully",
	}
}

func UpdateBook(c context.Context, id string, body io.ReadCloser) entities.BookChangesResponse {
	book := entities.Book{}
	json.NewDecoder(body).Decode(&book)

	for i, b := range repositories.Books {
		if b.ID == id {
			repositories.Books[i] = book
		}
	}

	return entities.BookChangesResponse{
		Status:  "success",
		Message: "Book updated successfully",
	}
}

func DeleteBook(c context.Context, id string) entities.BookChangesResponse {
	for _, b := range repositories.Books {
		if b.ID == id {
			repositories.Books = slices.DeleteFunc(repositories.Books, func(s entities.Book) bool {
				return s.ID == id
			})
		}
	}
	return entities.BookChangesResponse{
		Status:  "success",
		Message: "Book deleted successfully",
	}
}
