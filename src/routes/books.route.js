import { BooksHandler } from "../handlers/books.handler.js";

const booksHandler = new BooksHandler();

export const booksRoute = [
  {
    method: "GET",
    path: "/books",
    handler: booksHandler.getAllBooks,
  },
  {
    method: "GET",
    path: "/books/{bookId}",
    handler: booksHandler.getDetailBook,
  },
  {
    method: "POST",
    path: "/books",
    handler: booksHandler.saveNewBook,
  },
  {
    method: "PUT",
    path: "/books/{bookId}",
    handler: booksHandler.updateBook,
  },
  {
    method: "DELETE",
    path: "/books/{bookId}",
    handler: booksHandler.deleteBook,
  },
];
