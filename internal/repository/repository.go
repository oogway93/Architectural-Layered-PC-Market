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


