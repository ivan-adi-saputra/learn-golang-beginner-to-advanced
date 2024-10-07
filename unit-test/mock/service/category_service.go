package service

import (
	"errors"
	"golang-beginner-to-advanced/unit-test/mock/entity"
	"golang-beginner-to-advanced/unit-test/mock/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}