package helper

import (
	"github.com/jintoples/rest-desent/model/domain"
	"github.com/jintoples/rest-desent/model/web"
)

func ToBookResponse(book domain.Book) web.BookResponse {
	return web.BookResponse{
		Id:   book.Id,
		Name: book.Name,
	}
}
