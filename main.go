package main

import (
	"net/http"
	"salman/belajar-golang-restful-api/app"
	"salman/belajar-golang-restful-api/controller"
	_ "salman/belajar-golang-restful-api/controller"
	"salman/belajar-golang-restful-api/helper"
	"salman/belajar-golang-restful-api/middleware"
	"salman/belajar-golang-restful-api/repository"
	"salman/belajar-golang-restful-api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
