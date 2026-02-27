package main

import (
	"net/http"

	"github.com/jintoples/rest-desent/app"
	"github.com/jintoples/rest-desent/controller"
	"github.com/jintoples/rest-desent/helper"
	"github.com/jintoples/rest-desent/middleware"
	"github.com/jintoples/rest-desent/repository"
	"github.com/jintoples/rest-desent/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDb()
	validate := validator.New()
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository, db, validate)
	bookController := controller.NewBookController(bookService)
	router := app.NewRouter(bookController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
