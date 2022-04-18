//* Contoh MOCK => Aplikasi Query Ke Database
package repository

import (
	"golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoriRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoriRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	}else{
		category := arguments.Get(0).(entity.Category)
		return &category
	}



}