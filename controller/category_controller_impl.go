package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sumitroajiprabowo/go-restfull-api/helper"
	"github.com/sumitroajiprabowo/go-restfull-api/model/web"
	"github.com/sumitroajiprabowo/go-restfull-api/service"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

// Create a new CategoryControllerImpl Insert method to handle POST requests to /categories
func (c *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var request web.CategoryCreateRequest

	helper.ReadFromRequestBody(r, &request)

	response := c.CategoryService.Create(r.Context(), request)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, webResponse)
}

// Update a new CategoryControllerImpl Update method to handle PUT requests to /categories/:id
func (c *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var request web.CategoryUpdateRequest

	helper.ReadFromRequestBody(r, &request)

	categoryId := ps.ByName("categoryId")
	// id, err := helper.StringToInt64(categoryId)
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	// request.Id = id
	request.Id = int64(id)

	response := c.CategoryService.Update(r.Context(), request)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, webResponse)
}

// Delete a new CategoryControllerImpl Delete method to handle DELETE requests to /categories/:id
func (c *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	categoryId := ps.ByName("categoryId")
	// id, err := helper.StringToInt64(categoryId)
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	c.CategoryService.Delete(r.Context(), int64(id))

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

// FindById a new CategoryControllerImpl FindById method to handle GET requests to /categories/:id
func (c *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	categoryId := ps.ByName("categoryId")
	// id, err := helper.StringToInt64(categoryId)
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	response := c.CategoryService.FindById(r.Context(), int64(id))

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, webResponse)
}

// FindAll a new CategoryControllerImpl FindAll method to handle GET requests to /categories
func (c *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	categoryResponses := c.CategoryService.FindAll(r.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
