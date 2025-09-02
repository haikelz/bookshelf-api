package entities

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
