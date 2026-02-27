package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jintoples/rest-desent/helper"
	"github.com/jintoples/rest-desent/model/domain"
)

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (repository *BookRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	sql := "INSERT INTO book (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, sql, book.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	book.Id = int(id)
	return book
}

func (repository *BookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	sql := "UPDATE book SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, book.Name, book.Id)
	helper.PanicIfError(err)

	return book
}

func (repository *BookRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, book domain.Book) {
	sql := "DELETE FROM book WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, book.Id)
	helper.PanicIfError(err)

	return
}

func (repository *BookRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bookId int) (domain.Book, error) {
	sql := "SELECT id, name FROM book WHERE id = ?"
	rows, err := tx.QueryContext(ctx, sql, bookId)
	helper.PanicIfError(err)
	defer rows.Close()

	book := domain.Book{}
	if rows.Next() {
		err := rows.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err)

		return book, nil
	} else {
		return book, errors.New("book not found")
	}
}

func (repository *BookRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Book {
	sql := "SELECT id, name FROM book"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		book := domain.Book{}
		err := rows.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err)

		books = append(books, book)
	}
	return books
}
