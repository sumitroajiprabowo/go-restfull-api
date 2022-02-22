package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/sumitroajiprabowo/go-restfull-api/app"
	"github.com/sumitroajiprabowo/go-restfull-api/controller"
	"github.com/sumitroajiprabowo/go-restfull-api/helper"
	"github.com/sumitroajiprabowo/go-restfull-api/middleware"
	"github.com/sumitroajiprabowo/go-restfull-api/model/entity"
	"github.com/sumitroajiprabowo/go-restfull-api/repository"
	"github.com/sumitroajiprabowo/go-restfull-api/service"
)

func setupNewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:kmzway87aa@tcp(localhost:3306)/golang_restfull_api_test")
	helper.PanicIfError(err)

	// set pool settings
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

/*
Create setup router with db connection and controller
*/
func setupRouter(db *sql.DB) http.Handler {

	validate := validator.New() // create new validator

	//create new category repository with db connection and validate object
	categoryRepository := repository.NewCategoryRepository()

	// create new category service with category repository and validate object
	categoryService := service.NewCategoryService(categoryRepository, db, validate)

	// create new category controller with category service
	categoryController := controller.NewCategoryController(categoryService)

	// create new router with category controller and middleware
	router := app.NewRouter(categoryController)

	// create new auth middleware with router
	return middleware.NewAuthMiddleware(router) //return router with middleware

}

// Create function for truncate category table
func trancateCategory(db *sql.DB) {
	db.Exec("TRUNCATE TABLE category") //truncate table categories
}

//Create test create category with valid data
func TestCreateCategoryValidData(t *testing.T) {

	db := setupNewDB() //setup new db connection

	defer db.Close() //close db connection

	trancateCategory(db) //truncate table every test

	router := setupRouter(db) //setup router with db connection

	responseBody := strings.NewReader(`{"name":"test"}`) //create request body

	//create request with method, url, and body
	request, err := http.NewRequest("POST", "http://localhost:8080/api/categories", responseBody)
	helper.PanicIfError(err)

	request.Header.Set("Content-Type", "application/json") //set header
	request.Header.Set("X-API-Key", "Authorization")       //set header with key Authorization

	//recorder
	recorder := httptest.NewRecorder() //create recorder

	router.ServeHTTP(recorder, request) //serve request

	response := recorder.Result()             //get response
	assert.Equal(t, 200, response.StatusCode) //assert status code

	body, err := ioutil.ReadAll(response.Body) //read response body
	helper.PanicIfError(err)

	var responseData map[string]interface{} //  create response data with map type

	json.Unmarshal(body, &responseData) //unmarshal response body to response data

	fmt.Println(responseData) //print response data

	assert.Equal(t, 200, int(responseData["code"].(float64)))
	assert.Equal(t, "test", responseData["data"].(map[string]interface{})["name"])
	assert.Equal(t, "OK", responseData["status"].(string))
}

func TestCreateCategoryInvalidData(t *testing.T) {

	db := setupNewDB() //setup new db connection

	defer db.Close() //close db connection

	trancateCategory(db) //truncate table every test

	router := setupRouter(db) //setup router with db connection

	responseBody := strings.NewReader(`{"name":""}`) //create request body

	//create request with method, url, and body
	request, err := http.NewRequest("POST", "http://localhost:8080/api/categories", responseBody)
	helper.PanicIfError(err)

	request.Header.Set("Content-Type", "application/json") //set header
	request.Header.Set("X-API-Key", "Authorization")       //set header with key Authorization

	//recorder
	recorder := httptest.NewRecorder() //create recorder

	router.ServeHTTP(recorder, request) //serve request

	response := recorder.Result()             //get response
	assert.Equal(t, 400, response.StatusCode) //assert status code

	body, err := ioutil.ReadAll(response.Body) //read response body
	helper.PanicIfError(err)

	var responseData map[string]interface{} //  create response data with map type

	json.Unmarshal(body, &responseData) //unmarshal response body to response data

	fmt.Println(responseData) //print response data

	assert.Equal(t, 400, int(responseData["code"].(float64)))
	assert.Equal(t, "Bad Request", responseData["status"].(string))
}

func TestUpdateCategorySuccess(t *testing.T) {

	db := setupNewDB() //setup new db connection

	defer db.Close() //close db connection

	trancateCategory(db) //truncate table every test

	tx, _ := db.Begin()                                      // begin transaction
	categoryRepository := repository.NewCategoryRepository() //create new category repository

	//create new category with name test and description test and save to database
	category := categoryRepository.Create(context.Background(), tx, entity.Category{
		Name: "Category 1",
	})
	tx.Commit() //commit transaction

	router := setupRouter(db) //setup router with db connection

	requestBody := strings.NewReader(`{"name" : "Category 2"}`) //update new request body with json data

	//create request with method, url, and body
	request, err := http.NewRequest("PUT", "http://localhost:8080/api/categories/"+strconv.Itoa(int(category.Id)), requestBody)
	helper.PanicIfError(err)

	request.Header.Set("Content-Type", "application/json") //set header
	request.Header.Set("X-API-Key", "Authorization")       //set header with key Authorization

	//recorder
	recorder := httptest.NewRecorder() //create recorder

	router.ServeHTTP(recorder, request) //serve request

	response := recorder.Result()             //get response
	assert.Equal(t, 200, response.StatusCode) //assert status code

	body, err := ioutil.ReadAll(response.Body) //read response body
	helper.PanicIfError(err)

	var responseData map[string]interface{} //  create response data with map type

	json.Unmarshal(body, &responseData) //unmarshal response body to response data

	fmt.Println(responseData) //print response data

	assert.Equal(t, 200, int(responseData["code"].(float64)))
	assert.Equal(t, "OK", responseData["status"].(string))
	assert.Equal(t, "Category 2", responseData["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryNotFound(t *testing.T) {

	db := setupNewDB() //setup new db connection

	defer db.Close() //close db connection

	trancateCategory(db) //truncate table every test

	router := setupRouter(db) //setup router with db connection

	requestBody := strings.NewReader(`{"name" : "Category 2"}`) //update new request body with json data

	//create request with method, url, and body
	request, err := http.NewRequest("PUT", "http://localhost:8080/api/categories/404", requestBody)
	helper.PanicIfError(err)

	request.Header.Set("Content-Type", "application/json") //set header
	request.Header.Set("X-API-Key", "Authorization")       //set header with key Authorization

	//recorder
	recorder := httptest.NewRecorder() //create recorder

	router.ServeHTTP(recorder, request) //serve request

	response := recorder.Result()             //get response
	assert.Equal(t, 404, response.StatusCode) //assert status code

	body, err := ioutil.ReadAll(response.Body) //read response body
	helper.PanicIfError(err)

	var responseData map[string]interface{} //  create response data with map type

	json.Unmarshal(body, &responseData) //unmarshal response body to response data

	fmt.Println(responseData) //print response data

	assert.Equal(t, 404, int(responseData["code"].(float64)))
	assert.Equal(t, "Not Found", responseData["status"].(string))
}

func TestGetCategoryById(t *testing.T) {

	db := setupNewDB() //setup new db connection

	defer db.Close() //close db connection

	trancateCategory(db) //truncate table every test

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, entity.Category{
		Name: "Category 1",
	})
	tx.Commit()

	router := setupRouter(db) //setup router with db connection

	//create request with method, url, and body
	request, err := http.NewRequest("GET", "http://localhost:8080/api/categories/"+strconv.Itoa(int(category.Id)), nil)
	helper.PanicIfError(err)

	request.Header.Set("Content-Type", "application/json") //set header
	request.Header.Set("X-API-Key", "Authorization")       //set header with key Authorization

	//recorder
	recorder := httptest.NewRecorder() //create recorder

	router.ServeHTTP(recorder, request) //serve request

	response := recorder.Result()             //get response
	assert.Equal(t, 200, response.StatusCode) //assert status code

	body, err := ioutil.ReadAll(response.Body) //read response body
	helper.PanicIfError(err)

	var responseData map[string]interface{} //  create response data with map type

	json.Unmarshal(body, &responseData) //unmarshal response body to response data

	fmt.Println(responseData) //print response data

	assert.Equal(t, 200, int(responseData["code"].(float64)))
	assert.Equal(t, "OK", responseData["status"].(string))
	assert.Equal(t, "Category 1", responseData["data"].(map[string]interface{})["name"])
	assert.Equal(t, category.Id, int64(responseData["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Name, responseData["data"].(map[string]interface{})["name"])

}

func TestGetCategoryByIdNotFound(t *testing.T) {

	db := setupNewDB() //setup new db connection

	defer db.Close() //close db connection

	trancateCategory(db) //truncate table every test

	router := setupRouter(db) //setup router with db connection

	//create request with method, url, and body
	request, err := http.NewRequest("GET", "http://localhost:8080/api/categories/404", nil)
	helper.PanicIfError(err)

	request.Header.Set("Content-Type", "application/json") //set header
	request.Header.Set("X-API-Key", "Authorization")       //set header with key Authorization

	//recorder
	recorder := httptest.NewRecorder() //create recorder

	router.ServeHTTP(recorder, request) //serve request

	response := recorder.Result()             //get response
	assert.Equal(t, 404, response.StatusCode) //assert status code

	body, err := ioutil.ReadAll(response.Body) //read response body
	helper.PanicIfError(err)

	var responseData map[string]interface{} //  create response data with map type

	json.Unmarshal(body, &responseData) //unmarshal response body to response data

	fmt.Println(responseData) //print response data

	assert.Equal(t, 404, int(responseData["code"].(float64)))
	assert.Equal(t, "Not Found", responseData["status"].(string))
}

func TestGetAllCategory(t *testing.T) {

	db := setupNewDB() //setup new db connection

	defer db.Close() //close db connection

	trancateCategory(db) //truncate table every test

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category1 := categoryRepository.Create(context.Background(), tx, entity.Category{
		Name: "Category 1",
	})
	category2 := categoryRepository.Create(context.Background(), tx, entity.Category{
		Name: "Category 2",
	})
	tx.Commit()

	router := setupRouter(db) //setup router with db connection

	//create request with method, url, and body
	request, err := http.NewRequest("GET", "http://localhost:8080/api/categories", nil)
	helper.PanicIfError(err)

	request.Header.Set("Content-Type", "application/json") //set header
	request.Header.Set("X-API-Key", "Authorization")       //set header with key Authorization

	//recorder
	recorder := httptest.NewRecorder() //create recorder

	router.ServeHTTP(recorder, request) //serve request

	response := recorder.Result()             //get response
	assert.Equal(t, 200, response.StatusCode) //assert status code

	body, err := ioutil.ReadAll(response.Body) //read response body
	helper.PanicIfError(err)

	var responseData map[string]interface{} //  create response data with map type

	json.Unmarshal(body, &responseData) //unmarshal response body to response data

	fmt.Println(responseData) //print response data

	assert.Equal(t, 200, int(responseData["code"].(float64)))
	assert.Equal(t, "OK", responseData["status"].(string))
	assert.Equal(t, category1.Id, int64(responseData["data"].([]interface{})[0].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category1.Name, responseData["data"].([]interface{})[0].(map[string]interface{})["name"])
	assert.Equal(t, category2.Id, int64(responseData["data"].([]interface{})[1].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category2.Name, responseData["data"].([]interface{})[1].(map[string]interface{})["name"])
}

func TestDeleteCategorySuccess(t *testing.T) {

	db := setupNewDB() //setup new db connection

	defer db.Close() //close db connection

	trancateCategory(db) //truncate table every test

	tx, _ := db.Begin()                                      // begin transaction
	categoryRepository := repository.NewCategoryRepository() //create new category repository

	//create new category with name test and description test and save to database
	category := categoryRepository.Create(context.Background(), tx, entity.Category{
		Name: "Category 1",
	})
	tx.Commit() //commit transaction

	router := setupRouter(db) //setup router with db connection

	//create request with method, url, and body
	request, err := http.NewRequest("DELETE", "http://localhost:8080/api/categories/"+strconv.Itoa(int(category.Id)), nil)
	helper.PanicIfError(err)

	request.Header.Set("Content-Type", "application/json") //set header
	request.Header.Set("X-API-Key", "Authorization")       //set header with key Authorization

	//recorder
	recorder := httptest.NewRecorder() //create recorder

	router.ServeHTTP(recorder, request) //serve request

	response := recorder.Result()             //get response
	assert.Equal(t, 200, response.StatusCode) //assert status code

	body, err := ioutil.ReadAll(response.Body) //read response body
	helper.PanicIfError(err)

	var responseData map[string]interface{} //  create response data with map type

	json.Unmarshal(body, &responseData) //unmarshal response body to response data

	fmt.Println(responseData) //print response data

	assert.Equal(t, 200, int(responseData["code"].(float64)))
	assert.Equal(t, "OK", responseData["status"].(string))
}

func TestDeleteCategoryNotFound(t *testing.T) {

	db := setupNewDB() //setup new db connection

	defer db.Close() //close db connection

	trancateCategory(db) //truncate table every test

	router := setupRouter(db) //setup router with db connection

	//create request with method, url, and body
	request, err := http.NewRequest("DELETE", "http://localhost:8080/api/categories/404", nil)
	helper.PanicIfError(err)

	request.Header.Set("Content-Type", "application/json") //set header
	request.Header.Set("X-API-Key", "Authorization")       //set header with key Authorization

	//recorder
	recorder := httptest.NewRecorder() //create recorder

	router.ServeHTTP(recorder, request) //serve request

	response := recorder.Result()             //get response
	assert.Equal(t, 404, response.StatusCode) //assert status code

	body, err := ioutil.ReadAll(response.Body) //read response body
	helper.PanicIfError(err)

	var responseData map[string]interface{} //  create response data with map type

	json.Unmarshal(body, &responseData) //unmarshal response body to response data

	fmt.Println(responseData) //print response data

	assert.Equal(t, 404, int(responseData["code"].(float64)))
	assert.Equal(t, "Not Found", responseData["status"].(string))
}

func TestUnauthorized(t *testing.T) {

	db := setupNewDB() //setup new db connection

	defer db.Close() //close db connection

	trancateCategory(db) //truncate table every test

	router := setupRouter(db) //setup router with db connection

	//create request with method, url, and body
	request, err := http.NewRequest("GET", "http://localhost:8080/api/categories", nil)
	helper.PanicIfError(err)

	request.Header.Set("Content-Type", "application/json") //set header
	request.Header.Set("X-API-Key", "Wrong")               //set header with key Authorization

	//recorder
	recorder := httptest.NewRecorder() //create recorder

	router.ServeHTTP(recorder, request) //serve request

	response := recorder.Result()             //get response
	assert.Equal(t, 401, response.StatusCode) //assert status code

	body, err := ioutil.ReadAll(response.Body) //read response body
	helper.PanicIfError(err)

	var responseData map[string]interface{} //  create response data with map type

	json.Unmarshal(body, &responseData) //unmarshal response body to response data

	fmt.Println(responseData) //print response data

	assert.Equal(t, 401, int(responseData["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseData["status"].(string))

}
