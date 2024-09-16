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
  constructor() {
    this.fail = "fail";
    this.success = "success";
  }
}
