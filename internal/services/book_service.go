package services

import (
	"bookshelf/internal/entities"
	"bookshelf/internal/repositories"
	"slices"

	"github.com/google/uuid"
)

func GetBooks() entities.BooksResponse {
	return entities.BooksResponse{
		Status: "success",
		Data:   repositories.Books,
	}
}

func GetBookById(id string) entities.BookDetailResponse {
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

func CreateBook(book entities.Book) entities.BookChangesResponse {
	book.ID = uuid.New().String()
	repositories.Books = append(repositories.Books, book)

	return entities.BookChangesResponse{
		Status:  "success",
		Message: "Book created successfully",
	}
}

func UpdateBook(id string, book entities.Book) entities.BookChangesResponse {
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

func DeleteBook(id string) entities.BookChangesResponse {
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
