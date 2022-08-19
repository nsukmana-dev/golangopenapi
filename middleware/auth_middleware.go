package middleware

import (
	"golangopenapi/helper"
	"golangopenapi/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	if "TAMVAN DAN BERANI" == request.Header.Get("X-API-Key") {
		middleware.Handler.ServeHTTP(write, request)
	} else {
		write.Header().Set("Content-Type", "application/json")
		write.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(write, webResponse)
	}
}
