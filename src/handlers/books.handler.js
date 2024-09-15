import { Book, books } from "../configs/data.js";

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
    const findBook = books.find((item) => item.bookId === bookId);

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
    const bookId = req.params.bookId;

    return {
      status: "success",
      message: "Buku berhasil diperbarui",
    };
  }

  deleteBook(req) {
    const bookId = req.params.bookId;

    return {
      status: "success",
      message: "Buku berhasil dihapus",
    };
  }
}
