package exception

import (
	"net/http"

	"github.com/jintoples/rest-desent/helper"
	"github.com/jintoples/rest-desent/model/web"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	exception, ok := err.(NotFoundError)
	if ok {
		notFoundError(w, r, exception)
		return
	}

	exceptionValidate, okValidate := err.(validator.ValidationErrors)
	if okValidate {
		validationError(w, r, exceptionValidate)
		return
	}

	InternalServerError(w, r, err)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	webResponse := web.WebResponse{
		Code:   http.StatusNotFound,
		Status: "Not Found",
		Data:   err,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func validationError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "Bad Request",
		Data:   err,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func InternalServerError(w http.ResponseWriter, request *http.Request, err interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResponseBody(w, webResponse)
}
