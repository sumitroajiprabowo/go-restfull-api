// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/sumitroajiprabowo/go-restfull-api/app"
	"github.com/sumitroajiprabowo/go-restfull-api/controller"
	"github.com/sumitroajiprabowo/go-restfull-api/middleware"
	"github.com/sumitroajiprabowo/go-restfull-api/repository"
	"github.com/sumitroajiprabowo/go-restfull-api/service"
	"net/http"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func InizializedServer() *http.Server {
	categoryRepositoryImpl := repository.NewCategoryRepository()
	db := app.NewDB()
	validate := validator.New()
	categoryServiceImpl := service.NewCategoryService(categoryRepositoryImpl, db, validate)
	categoryControllerImpl := controller.NewCategoryController(categoryServiceImpl)
	router := app.NewRouter(categoryControllerImpl)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)
	return server
}

// wire.go:

var categorySet = wire.NewSet(repository.NewCategoryRepository, wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)), service.NewCategoryService, wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)), controller.NewCategoryController, wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)))
