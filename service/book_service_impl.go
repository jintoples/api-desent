package service

import (
	"context"
	"database/sql"

	"github.com/jintoples/rest-desent/exception"
	"github.com/jintoples/rest-desent/helper"
	"github.com/jintoples/rest-desent/model/domain"
	"github.com/jintoples/rest-desent/model/web"
	"github.com/jintoples/rest-desent/repository"

	"github.com/go-playground/validator/v10"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewBookService(bookService repository.BookRepository, DB *sql.DB, validate *validator.Validate) BookService {
	return &BookServiceImpl{
		BookRepository: bookService,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *BookServiceImpl) Save(ctx context.Context, request web.BookCreateRequest) web.BookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer func() {
		helper.CommitOrRollback(tx)
	}()

	book := domain.Book{
		Name: request.Name,
	}

	book = service.BookRepository.Save(ctx, tx, book)

	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) Update(ctx context.Context, request web.BookUpdateRequest) web.BookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer func() {
		helper.CommitOrRollback(tx)
	}()

	book, err := service.BookRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	book.Name = request.Name

	book = service.BookRepository.Update(ctx, tx, book)

	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) Delete(ctx context.Context, bookId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer func() {
		helper.CommitOrRollback(tx)
	}()

	book, err := service.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.BookRepository.Delete(ctx, tx, book)
}

func (service *BookServiceImpl) FindById(ctx context.Context, bookId int) (web.BookResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer func() {
		helper.CommitOrRollback(tx)
	}()

	book, err := service.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToBookResponse(book), nil
}

func (service *BookServiceImpl) FindAll(ctx context.Context) []web.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer func() {
		helper.CommitOrRollback(tx)
	}()

	books := service.BookRepository.FindAll(ctx, tx)

	var bookResponses []web.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, helper.ToBookResponse(book))
	}

	return bookResponses
}
