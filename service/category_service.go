package service

import (
	"context"

	"github.com/jintoples/rest-desent/model/web"
)

type BookService interface {
	Save(ctx context.Context, request web.BookCreateRequest) web.BookResponse
	Update(ctx context.Context, request web.BookUpdateRequest) web.BookResponse
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) (web.BookResponse, error)
	FindAll(ctx context.Context) []web.BookResponse
}
