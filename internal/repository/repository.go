package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/shop/repositoryShop"
)

type CategoryRepository interface {
	GetAll() ([]products.Category,error)
}

type Repository struct {
	CategoryRepository
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		CategoryRepository: repositoryShop.NewCategoryShop(db),
	}
}
