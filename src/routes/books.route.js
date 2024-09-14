export const booksRoute = [
  {
    method: "GET",
    path: "/",
    handler: () => {
      return {
        statusCode: 200,
        message: "Success get data!",
      };
    },
  },
  {
    method: "GET",
    path: "/books/:bookId",
    handler: () => {
      /*
      return {
        status: "fail",
        message: "Buku tidak ditemukan",
      };*/
      return {
        status: "success",
        data: {
          book: {
            id: "aWZBUW3JN_VBE-9I",
            name: "Buku A Revisi",
            year: 2011,
            author: "Jane Doe",
            summary: "Lorem Dolor sit Amet",
            publisher: "Dicoding",
            pageCount: 200,
            readPage: 26,
            finished: false,
            reading: false,
            insertedAt: "2021-03-05T06:14:28.930Z",
            updatedAt: "2021-03-05T06:14:30.718Z",
          },
        },
      };
    },
  },
  {
    method: "POST",
    path: "/books",
    handler: () => {
      return {
        status: "success",
        data: {
          books: [
            {
              id: "Qbax5Oy7L8WKf74l",
              name: "Buku A",
              publisher: "Dicoding Indonesia",
            },
            {
              id: "1L7ZtDUFeGs7VlEt",
              name: "Buku B",
              publisher: "Dicoding Indonesia",
            },
            {
              id: "K8DZbfI-t3LrY7lD",
              name: "Buku C",
              publisher: "Dicoding Indonesia",
            },
          ],
        },
      };
    },
  },
  {
    method: "PUT",
    path: "/books/:bookId",
    handler: () => {
      return {
        status: "success",
        message: "Buku berhasil diperbarui",
      };
    },
  },
  {
    method: "DELETE",
    path: "/books/:bookId",
    handler: () => {
      return {
        status: "success",
        message: "Buku berhasil dihapus",
      };
    },
  },
];
