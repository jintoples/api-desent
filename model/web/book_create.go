package web

type BookCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type BookResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
