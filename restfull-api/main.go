package main

import (
	"golang-beginner-to-advanced/restfull-api/app"
	_ "github.com/go-sql-driver/mysql"
	"golang-beginner-to-advanced/restfull-api/controller"
	"golang-beginner-to-advanced/restfull-api/helper"
	"golang-beginner-to-advanced/restfull-api/middleware"
	"golang-beginner-to-advanced/restfull-api/repository"
	"golang-beginner-to-advanced/restfull-api/service"
	"net/http"

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