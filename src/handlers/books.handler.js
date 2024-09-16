import { AboutProject, Status } from "../entities/books.entity.js";
import { books } from "../repositories/books.repository.js";
import { escapeHtml } from "@hapi/hoek";
import { nanoid } from "nanoid";

export class BooksHandler {
  getHome(_, h) {
    const aboutProject = new AboutProject();

    return h
      .response({
        status: Status.success,
        message: "Sukses mendapatkan data About this project!",
        data: aboutProject,
      })
      .code(200);
  }

  getAllBooks(req, h) {
    const status = new Status();
    const query = req.query;

    if (query.name || query.reading || query.finished) {
      const filteredBooks = books.filter((item) => {
        let isFiltered = true;

        for (const key in query) {
          isFiltered =
            isFiltered &&
            item[key] ==
              (key === "reading" || key === "finished"
                ? Number(escapeHtml(query[key]))
                : escapeHtml(query[key]));
        }

        return isFiltered;
      });

      return h
        .response({
          status: status.success,
          data: {
            books: filteredBooks,
          },
        })
        .code(200);
    } else {
      const selectedDataBooks = books.map((item) => ({
        id: item.id,
        name: item.name,
        publisher: item.publisher,
      }));

      return h
        .response({
          status: status.success,
          data: {
            books: selectedDataBooks,
          },
        })
        .code(200);
    }
  }

  getDetailBook(req, h) {
    const status = new Status();
    const bookId = escapeHtml(req.params.bookId);
    const findBook = books.find((item) => item.id === bookId);

    if (findBook && bookId) {
      return h
        .response({
          status: status.success,
          data: {
            book: findBook,
          },
        })
        .code(200);
    } else {
      return h
        .response({
          status: status.fail,
          message: "Buku tidak ditemukan",
        })
        .code(404);
    }
  }

  saveNewBook(req, h) {
    const status = new Status();
    const { name, pageCount, readPage } = req.payload;

    if (!name) {
      return h
        .response({
          status: status.fail,
          message: "Gagal menambahkan buku. Mohon isi nama buku",
        })
        .code(400);
    }

    if (readPage > pageCount) {
      return h
        .response({
          status: status.fail,
          message:
            "Gagal menambahkan buku. readPage tidak boleh lebih besar dari pageCount",
        })
        .code(400);
    }

    const dateNow = new Date().toISOString();
    const bookId = nanoid(16);

    books.push({
      id: bookId,
      ...req.payload,
      insertedAt: dateNow,
      updatedAt: dateNow,
    });

    return h
      .response({
        status: status.success,
        message: "Buku berhasil ditambahkan",
        data: {
          bookId,
        },
      })
      .code(201);
  }

  updateBook(req, h) {
    const status = new Status();
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
          status: status.fail,
          message: "Gagal memperbarui buku. Mohon isi nama buku",
        })
        .code(400);
    }

    if (readPage > pageCount) {
      return h
        .response({
          status: status.fail,
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
          status: status.success,
          message: "Buku berhasil diperbarui",
        })
        .code(200);
    } else {
      return h
        .response({
          status: status.fail,
          message: "Gagal memperbarui buku. Id tidak ditemukan",
        })
        .code(404);
    }
  }

  deleteBook(req, h) {
    const status = new Status();
    const bookId = req.params.bookId;
    const findSameId = books.findIndex((item) => item.id === bookId);

    if (findSameId !== -1) {
      books.splice(findSameId, 1);

      return h
        .response({
          status: status.success,
          message: "Buku berhasil dihapus",
        })
        .code(200);
    }
  }
}
