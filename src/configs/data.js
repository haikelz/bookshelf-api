export class Book {
  constructor(
    name,
    year,
    author,
    summary,
    publisher,
    pageCount,
    readPage,
    reading
  ) {
    this.name = name;
    this.year = year;
    this.author = author;
    this.summary = summary;
    this.publisher = publisher;
    this.pageCount = pageCount;
    this.readPage = readPage;
    this.reading = reading;
  }
}

export let books = [];

export class AboutProject {
  constructor(
    author = "Haikel Ilham Hakim",
    repository = "https://github.com/haikelz/bookshelf-api"
  ) {
    this.author = author;
    this.repository = repository;
  }
}

export class Status {
  constructor(fail = "fail", success = "success") {
    this.fail = fail;
    this.success = success;
  }
}
