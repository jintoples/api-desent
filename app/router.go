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

	router.GET("/echo", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		webResponse := web.GeneralResponse{
			Success: true,
		}
		helper.WriteToResponseBody(w, webResponse)
	})

	router.GET("/books", bookController.FindAll)
	router.POST("/books", bookController.Create)
	router.GET("/books/:id", bookController.FindById)
	router.PUT("/books/:id", bookController.Update)
	router.DELETE("/books/:id", bookController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
