package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Create(writer http.ResponseWriter, request http.Request, params httprouter.Params)
	Update()
	Delete()
	FindById()
	FindAll()
}
