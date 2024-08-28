package repositoryPostgresShop

import (
	"log"
	"time"

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

func (d *CategoryShopPostgres) Create(newCategory models.Category) {
	tx := d.db.Begin()
	categoryName := newCategory.CategoryName

	var existingCategory models.Category
	result := d.db.Unscoped().Where("category_name = ? AND deleted_at IS NOT NULL", categoryName).First(&existingCategory)

	if result.Error == nil {
		rawSQL := `UPDATE categories SET deleted_at = NULL WHERE category_name = ?`
		d.db.Exec(rawSQL, existingCategory.CategoryName)
		log.Printf("Restored category: %s", categoryName)
	} else {
		result = d.db.Create(&newCategory)
		if result.Error != nil {
			log.Printf("Error creating new category: %v", result.Error)
		} else {
			log.Printf("Created new category: %s", categoryName)
		}
	}
	tx.Commit()
}

func (d *CategoryShopPostgres) GetAll() []map[string]interface{} {
	var categories []models.Category
	tx := d.db.Begin()
	result := d.db.Find(&categories)

	if result.Error != nil {
		log.Printf("Error finding records from category: %v", result.Error)
	}
	var resultCategories []map[string]interface{}
	for _, category := range categories {
		resultCategories = append(resultCategories, map[string]interface{}{
			"categoryName": category.CategoryName,
		})
	}
	tx.Commit()
	return resultCategories
}

func (d *CategoryShopPostgres) Delete(categoryID string) error {
	var category models.Category
	tx := d.db.Begin()
	result := d.db.Where("category_name = ?", categoryID).Delete(&category)
	if result.Error != nil {
		return result.Error
	}
	tx.Commit()
	return result.Error
}

func (d *CategoryShopPostgres) Get(categoryID string) string {
	var category models.Category
	result := d.db.Where("category_name = ?", categoryID).First(&category)
	if result.Error != nil {
		log.Fatalf("Error repositrory: %v", result)
	}
	return category.CategoryName
}

func (d *CategoryShopPostgres) Update(categoryID string, newCategory models.Category) error {
	var category models.Category
	tx := d.db.Begin()
	result := d.db.Where("category_name = ? AND deleted_at IS NULL", categoryID).First(&category)
	if result.Error != nil {
		return result.Error
	}
	category.CategoryName = newCategory.CategoryName
	category.UpdatedAt = time.Now()
	result = d.db.Save(&category)
	tx.Commit()
	return result.Error
}
