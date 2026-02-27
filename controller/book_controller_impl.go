package controller

import (
	"net/http"
	"strconv"

	"github.com/jintoples/rest-desent/helper"
	"github.com/jintoples/rest-desent/model/web"
	"github.com/jintoples/rest-desent/service"

	"github.com/julienschmidt/httprouter"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &BookControllerImpl{
		BookService: bookService,
	}
}

func (controller *BookControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	request := web.BookCreateRequest{}
	helper.ReadFromRequestBody(r, &request)

	bookResponse := controller.BookService.Save(r.Context(), request)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *BookControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	request := web.BookUpdateRequest{}
	helper.ReadFromRequestBody(r, &request)

	id := p.ByName("id")
	bookId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	request.Id = bookId

	bookResponse := controller.BookService.Update(r.Context(), request)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *BookControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	bookId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	controller.BookService.Delete(r.Context(), bookId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *BookControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	bookId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	bookResponse, err := controller.BookService.FindById(r.Context(), bookId)
	helper.PanicIfError(err)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *BookControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookResponse := controller.BookService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
