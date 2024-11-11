package helper

import (
	"golang-beginner-to-advanced/restfull-api/model/domain"
	"golang-beginner-to-advanced/restfull-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
        Id:   category.Id,
        Name: category.Name,
    }
}