package repositoryPostgresShop

import (
	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"gorm.io/gorm"
)

type CategoryShopPostgres struct {
	db *gorm.DB
}

func NewRepositoryCategoryShop(db *gorm.DB) *CategoryShopPostgres {
	return &CategoryShopPostgres{
		db: db,
	}
}

func (d *CategoryShopPostgres) GetAll() ([]products.Category, error) {
	return nil, nil
}
