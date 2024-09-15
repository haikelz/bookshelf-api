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

export let books = [
  {
    id: "1a2b3c",
    name: "The Art of Programming",
    year: 2020,
    author: "John Doe",
    summary:
      "A comprehensive guide to modern programming practices and patterns.",
    publisher: "Tech Books Publishing",
    pageCount: 500,
    readPage: 500,
    finished: true,
    reading: false,
    insertedAt: "2022-03-01T10:30:00Z",
    updatedAt: "2022-05-01T12:00:00Z",
  },
  {
    id: "4d5e6f",
    name: "Learning JavaScript",
    year: 2019,
    author: "Jane Smith",
    summary:
      "An introduction to JavaScript for beginners with practical examples.",
    publisher: "Code Masters Press",
    pageCount: 300,
    readPage: 120,
    finished: false,
    reading: true,
    insertedAt: "2022-04-15T09:45:00Z",
    updatedAt: "2022-04-16T11:00:00Z",
  },
  {
    id: "7g8h9i",
    name: "Mastering Data Structures",
    year: 2018,
    author: "Alice Brown",
    summary:
      "Deep dive into data structures and algorithms with real-world examples.",
    publisher: "Algorithm Press",
    pageCount: 600,
    readPage: 600,
    finished: true,
    reading: false,
    insertedAt: "2021-12-10T14:20:00Z",
    updatedAt: "2022-01-05T16:30:00Z",
  },
  {
    id: "10j11k12",
    name: "React for Professionals",
    year: 2021,
    author: "Michael Scott",
    summary:
      "Advanced React techniques and best practices for building dynamic web apps.",
    publisher: "Web Dev Books",
    pageCount: 450,
    readPage: 300,
    finished: false,
    reading: true,
    insertedAt: "2023-01-20T08:10:00Z",
    updatedAt: "2023-02-15T10:50:00Z",
  },
  {
    id: "13l14m15",
    name: "Python for Data Science",
    year: 2017,
    author: "Sara Connor",
    summary:
      "A practical approach to learning Python and using it for data analysis and machine learning.",
    publisher: "Data Science Publishers",
    pageCount: 350,
    readPage: 350,
    finished: true,
    reading: false,
    insertedAt: "2020-08-05T17:00:00Z",
    updatedAt: "2020-09-01T19:00:00Z",
  },
];

export const aboutProject = {
  author: "Haikel Ilham Hakim",
  repository: "https://github.com/haikelz/bookshelf-api",
};
