package main

import (
	"golangopenapi/app"
	"golangopenapi/controller"
	"golangopenapi/repository"
	"golangopenapi/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	CategoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", CategoryController.FindAll)
	router.GET("/api/categories/:categoryId", CategoryController.FindById)
	router.POST("/api/categories", CategoryController.Create)
	router.PUT("/api/categories", CategoryController.Update)
	router.DELETE("/api/categories", CategoryController.Delete)
}
