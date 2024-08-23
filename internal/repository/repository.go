package repository

import (
	"github.com/oogway93/golangArchitecture/internal/entity/products"
)


type CategoryRepository interface {
	GetAll() ([]products.Category,error)
}

type Repository struct {
	CategoryRepository CategoryRepository
}

// func NewRepository(db *sqlx.DB) *Repository{
// 	return &Repository{
// 		CategoryRepository: NewCategoryShop(),
// 	}
// }
