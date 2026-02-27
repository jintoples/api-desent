package middleware

import (
	"net/http"

	"github.com/jintoples/rest-desent/helper"
	"github.com/jintoples/rest-desent/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("X-API-Key")
	if "API_KEY" == key {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}
