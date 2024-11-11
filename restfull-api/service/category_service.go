package service

import (
	"context"
	"golang-beginner-to-advanced/restfull-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, requet web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, requet web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}