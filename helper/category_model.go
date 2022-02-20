package helper

import (
	"github.com/sumitroajiprabowo/go-restfull-api/model/entity"
	"github.com/sumitroajiprabowo/go-restfull-api/model/web"
)

func ToCategoryResponse(category *entity.Category) *web.CategoryResponse {
	return &web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

// func ToCategoryResponse(category *entity.Category) web.CategoryResponse {
// 	return web.CategoryResponse{
// 		Id:   category.Id,
// 		Name: category.Name,
// 	}
// }

func ToCategoriesResponse(categories []entity.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, *ToCategoryResponse(&category))
	}
	return categoryResponses
}
