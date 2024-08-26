package repositoryPostgresShop

import (
	"log"

	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
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

func (d *CategoryShopPostgres) Create(newCategory  models.Category) {
	tx := d.db.Begin()

	result := d.db.Create(&newCategory)

	if result.Error != nil {
		log.Printf("Error creating new category: %v", result.Error)
	}
	tx.Commit()
}
