package app

import (
	"net/http"

	"github.com/jintoples/rest-desent/controller"
	"github.com/jintoples/rest-desent/exception"
	"github.com/jintoples/rest-desent/helper"
	"github.com/jintoples/rest-desent/model/web"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(bookController controller.BookController) http.Handler {
	router := httprouter.New()

	router.GET("/ping/ping", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		webResponse := web.GeneralResponse{
			Success: true,
		}
		helper.WriteToResponseBody(w, webResponse)
	})

	router.POST("/ping/echo", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		helper.WriteToResponseBody(w, map[string]string{})
	})

	router.GET("/ping/books", bookController.FindAll)
	router.POST("/ping/books", bookController.Create)
	router.GET("/ping/books/:id", bookController.FindById)
	router.PUT("/ping/books/:id", bookController.Update)
	router.DELETE("/ping/books/:id", bookController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
