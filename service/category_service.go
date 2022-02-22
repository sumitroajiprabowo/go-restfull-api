package service

import (
	"context"

	"github.com/sumitroajiprabowo/go-restfull-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int64)
	FindById(ctx context.Context, categoryId int64) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
