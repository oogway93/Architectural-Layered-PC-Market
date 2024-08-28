package repositoryPostgresShop

import (
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
	"gorm.io/gorm"
)

type ProductShopPostgres struct {
	db *gorm.DB
}

func NewRepositoryProductShop(db *gorm.DB) *ProductShopPostgres {
	return &ProductShopPostgres{
		db: db,
	}
}

func (d *ProductShopPostgres) Create(requestData models.Product)
func (d *ProductShopPostgres) GetAll() []map[string]interface{}
func (d *ProductShopPostgres) Delete(categoryID string, productID string) error
func (d *ProductShopPostgres) Get(categoryID string, productID string) string
func (d *ProductShopPostgres) Update(categoryID string, productID string, newCategory models.Product) error
