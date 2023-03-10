package helper

import (
	"naufalatha/golang-restapi/model/domain"
	"naufalatha/golang-restapi/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var responses []web.CategoryResponse
	for _, category := range categories {
		responses = append(responses, ToCategoryResponse(category))
	}
	return responses
}
