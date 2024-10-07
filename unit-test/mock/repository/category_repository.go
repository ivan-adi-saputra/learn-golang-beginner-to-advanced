package repository

import "golang-beginner-to-advanced/unit-test/mock/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}