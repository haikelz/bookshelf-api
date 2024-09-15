import { STATUS } from "../configs/constants.js";
import { Book, books } from "../configs/data.js";
import { escapeHtml } from "@hapi/hoek";
import { nanoid } from "nanoid";

export class BooksHandler {
  book;

  constructor() {
    this.book = new Book();
  }

  getAllBooks(req, h) {
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

    books.push({ id: nanoid(), ...req.payload });

    return h
      .response({
        status: STATUS.success,
        message: "Buku berhasil ditambahkan",
      })
      .code(201);
  }

  updateBook(req, h) {
    const bookId = escapeHtml(req.params.bookId);
    // books.map((item) => (item.id === bookId ? { ...req.payload } : item));

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

    // 400
    if (!name) {
      return h
        .response({
          status: STATUS.fail,
          message: "Gagal memperbarui buku. Mohon isi nama buku",
        })
        .code(400);
    }

    // 400
    if (readPage > pageCount) {
      return h
        .response({
          status: STATUS.fail,
          message:
            "Gagal memperbarui buku. readPage tidak boleh lebih besar dari pageCount",
        })
        .code(400);
    }

    // 404
    if (!bookId) {
      return h
        .response({
          status: STATUS.fail,
          message: "Gagal memperbarui buku. Id tidak ditemukan",
        })
        .code(404);
    }

    return h
      .response({
        status: STATUS.success,
        message: "Buku berhasil diperbarui",
      })
      .code(200);
  }

  deleteBook(req, h) {
    const bookId = req.params.bookId;
    const findSameId = books.findIndex((item) => item.id === bookId);

    if (findSameId === -1) {
      books.splice(findSameId, 1);
    }

    return h
      .response({
        status: STATUS.success,
        message: "Buku berhasil dihapus",
      })
      .code(201);
  }
}
