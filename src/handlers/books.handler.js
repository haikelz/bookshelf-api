import { Book, books } from "../configs/data.js";
import { boomify } from "@hapi/boom";
import { escapeHtml } from "@hapi/hoek";
import { nanoid } from "nanoid";

export class BooksHandler {
  book;

  constructor() {
    this.book = new Book();
  }

  getAllBooks() {
    return {
      status: "success",
      data: {
        books,
      },
    };
  }

  getDetailBook(req) {
    const bookId = escapeHtml(req.params.bookId);
    const findBook = books.find((item) => item.id === bookId);

    if (findBook) {
      return {
        status: "success",
        data: {
          book: findBook,
        },
      };
    } else {
      const error = new Error("Buku tidak ditemukan");

      return boomify(error, {
        statusCode: 404,
      });
    }
  }

  saveNewBook(req) {
    books.push({ id: nanoid(), ...req.payload });

    return {
      status: "success",
      message: "Buku berhasil ditambahkan",
    };
  }

  updateBook(req) {
    const bookId = escapeHtml(req.params.bookId);
    books.map((item) => (item.id === bookId ? { ...req.payload } : item));

    return {
      status: "success",
      message: "Buku berhasil diperbarui",
    };
  }

  deleteBook(req) {
    const bookId = req.params.bookId;

    const findSameId = books.findIndex((item) => item.id === bookId);

    if (findSameId === -1) {
      books.splice(findSameId, 1);
    }

    return {
      status: "success",
      message: "Buku berhasil dihapus",
    };
  }
}
