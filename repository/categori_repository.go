//* Contoh MOCK => Aplikasi Query Ke Database
package repository

import "golang-unit-test/entity"

type CategoriRepository interface {
	FindById(id string) *entity.Category
}