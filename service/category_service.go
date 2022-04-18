//* Contoh MOCK => Aplikasi Query Ke Database
package service

import (
	"errors"
	"golang-unit-test/entity"
	"golang-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoriRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("category not found")
	}else{
		return category, nil
	}
}