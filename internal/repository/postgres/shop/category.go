package repositoryPostgresShop

import (
	"github.com/jmoiron/sqlx"
	"github.com/oogway93/golangArchitecture/internal/entity/products"
)

type CategoryShopPostgres struct {
	db *sqlx.DB
}
func NewRepositoryCategoryShop(db *sqlx.DB) *CategoryShopPostgres {
	return &CategoryShopPostgres{
		db: db,
	}
}

func (d *CategoryShopPostgres) GetAll() ([]products.Category, error) {
	return nil, nil
}
