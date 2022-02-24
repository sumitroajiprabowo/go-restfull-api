package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/sumitroajiprabowo/go-restfull-api/exception"
	"github.com/sumitroajiprabowo/go-restfull-api/helper"
	"github.com/sumitroajiprabowo/go-restfull-api/model/entity"
	"github.com/sumitroajiprabowo/go-restfull-api/model/web"
	"github.com/sumitroajiprabowo/go-restfull-api/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

// CategoryServiceImplementation.Create with Rollback
func (c *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {

	//validate request
	err := c.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := c.DB.BeginTx(ctx, nil) // begin transaction
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx, err) // commit or rollback

	category := c.CategoryRepository.Create(ctx, tx, entity.Category{
		Name: request.Name,
	})

	return *helper.ToCategoryResponse(&category)
}

// CategoryServiceImplementation.Update with Rollback and check if category exist and not found
func (c *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {

	//validate request
	err := c.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := c.DB.BeginTx(ctx, nil) // begin transaction
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx, err) // commit or rollback

	// check if category exist and not found
	category, err := c.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// update category with new name
	category.Name = request.Name

	// update category
	category = c.CategoryRepository.Update(ctx, tx, category)

	return *helper.ToCategoryResponse(&category)

}

// CategoryServiceImplementation.Delete with Rollback and check if category not found
func (c *CategoryServiceImpl) Delete(ctx context.Context, categoryId int64) {

	tx, err := c.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx, err)

	// check if category not found
	category, err := c.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// delete category
	c.CategoryRepository.Delete(ctx, tx, category)
}

// CategoryServiceImplementation.FindById with Rollback and check if category not found
func (c *CategoryServiceImpl) FindById(ctx context.Context, categoryId int64) web.CategoryResponse {

	tx, err := c.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx, err) // commit or rollback

	// check if category not found
	category, err := c.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return *helper.ToCategoryResponse(&category)
}

// CategoryServiceImplementation.FindAll with Rollback
func (c *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {

	tx, err := c.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx, err)

	categories := c.CategoryRepository.FindAll(ctx, tx)

	// categoriesResponse := []web.CategoryResponse{}

	// for i, category := range categories {
	// 	categoriesResponse[i] = *helper.ToCategoryResponse(&category)
	// }

	// for _, category := range categories {
	// 	categoriesResponse = append(categoriesResponse, *helper.ToCategoryResponse(&category))
	// }

	// return categoriesResponse

	return helper.ToCategoriesResponse(categories)

}
