package entities

type Book struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Year      string `json:"year"`
	Author    string `json:"author"`
	Summary   string `json:"summary"`
	Publisher string `json:"publisher"`
	PageCount int    `json:"pageCount"`
	ReadPage  int    `json:"readPage"`
	Reading   bool   `json:"reading"`
}

type BooksResponse struct {
	Status string `json:"status"`
	Data   []Book `json:"data"`
}

type BookDetailResponse struct {
	Status string `json:"status"`
	Data   Book   `json:"data"`
}

type BookChangesResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
