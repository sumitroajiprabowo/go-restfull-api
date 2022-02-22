package middleware

import (
	"net/http"

	"github.com/sumitroajiprabowo/go-restfull-api/helper"
	"github.com/sumitroajiprabowo/go-restfull-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (am *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if "Authorization" == r.Header.Get("X-API-Key") {
		am.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(w, webResponse)

	}

}
