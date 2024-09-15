import { STATUS } from "../configs/constants.js";
import { Book, books } from "../configs/data.js";
import { escapeHtml } from "@hapi/hoek";
import { nanoid } from "nanoid";

export class BooksHandler {
  book;

  constructor() {
    this.book = new Book();
  }

  getAllBooks(_, h) {
    return h
      .response({
        status: STATUS.success,
        data: {
          books,
        },
      })
      .code(200);
  }

  getDetailBook(req, h) {
    const bookId = escapeHtml(req.params.bookId);
    const findBook = books.find((item) => item.id === bookId);

    if (findBook) {
      return h
        .response({
          status: STATUS.success,
          data: {
            book: findBook,
          },
        })
        .code(200);
    } else {
      return h
        .response({
          status: STATUS.fail,
          message: "Buku tidak ditemukan",
        })
        .code(404);
    }
  }

  saveNewBook(req, h) {
    const { name, pageCount, readPage } = req.payload;

    if (!name) {
      return h
        .response({
          status: STATUS.fail,
          message: "Gagal menambahkan buku. Mohon isi nama buku",
        })
        .code(400);
    }

    if (readPage > pageCount) {
      return h
        .response({
          status: STATUS.fail,
          message:
            "Gagal menambahkan buku. readPage tidak boleh lebih besar dari pageCount",
        })
        .code(400);
    }

    books.push({
      id: nanoid(),
      ...req.payload,
      insertedAt: new Date(),
      updatedAt: null,
    });

    return h
      .response({
        status: STATUS.success,
        message: "Buku berhasil ditambahkan",
      })
      .code(201);
  }

  updateBook(req, h) {
    const bookId = escapeHtml(req.params.bookId);

    const {
      name,
      year,
      author,
      summary,
      publisher,
      pageCount,
      readPage,
      reading,
    } = req.payload;

    const findIndex = books.findIndex((item) => item.id === bookId);

    if (!name) {
      return h
        .response({
          status: STATUS.fail,
          message: "Gagal memperbarui buku. Mohon isi nama buku",
        })
        .code(400);
    }

    if (readPage > pageCount) {
      return h
        .response({
          status: STATUS.fail,
          message:
            "Gagal memperbarui buku. readPage tidak boleh lebih besar dari pageCount",
        })
        .code(400);
    }

    if (bookId && findIndex !== -1) {
      books[findIndex] = {
        ...books[findIndex],
        name,
        year,
        author,
        summary,
        publisher,
        pageCount,
        readPage,
        reading,
        updatedAt: new Date().toISOString(),
      };

      return h
        .response({
          status: STATUS.success,
          message: "Buku berhasil diperbarui",
        })
        .code(200);
    } else {
      return h
        .response({
          status: STATUS.fail,
          message: "Gagal memperbarui buku. Id tidak ditemukan",
        })
        .code(400);
    }
  }

  deleteBook(req, h) {
    const bookId = req.params.bookId;
    const findSameId = books.findIndex((item) => item.id === bookId);

    if (findSameId !== -1) {
      books.splice(findSameId, 1);

      return h
        .response({
          status: STATUS.success,
          message: "Buku berhasil dihapus",
        })
        .code(201);
    }
  }
}
