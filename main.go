package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sumitroajiprabowo/go-restfull-api/helper"
	"github.com/sumitroajiprabowo/go-restfull-api/middleware"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: authMiddleware,
	}
}

func main() {

	/*
		Exmaple Manual
		not use Dependency Injection
	*/

	// db := app.NewDB()
	// validate := validator.New()
	// categoryRepository := repository.NewCategoryRepository()
	// categoryService := service.NewCategoryService(categoryRepository, db, validate)
	// categoryController := controller.NewCategoryController(categoryService)

	// router := app.NewRouter(categoryController)
	// authMiddleware := middleware.NewAuthMiddleware(router)
	// server := NewServer(authMiddleware)

	server := InizializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
