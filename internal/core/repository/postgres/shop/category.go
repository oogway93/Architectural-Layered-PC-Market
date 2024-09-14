package repositoryPostgresShop

import (
	"log/slog"
	"time"

	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
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

func (d *CategoryShopPostgres) Create(newCategory *models.Category) {
	tx := d.db.Begin()
	categoryName := newCategory.CategoryName

	var existingCategory models.Category
	result := tx.Unscoped().Where("category_name = ?", categoryName).First(&existingCategory)

	if result.Error == nil {
		rawSQL := `UPDATE categories SET deleted_at = NULL WHERE category_name = ?`
		tx.Exec(rawSQL, existingCategory.CategoryName)
		slog.Info("Restored category", "categoryName", categoryName)
	} else {
		result = tx.Create(&newCategory)
		if result.Error != nil {
			slog.Warn("Error creating NEW CATEGORY", "error", result.Error)
		} else {
			slog.Info("Created NEW CATEGORY", "categoryName", categoryName)
		}
	}
	tx.Commit()
}

func (d *CategoryShopPostgres) GetAll() []map[string]interface{} {
	var categories []models.Category
	var products []models.Product
	tx := d.db.Begin()

	result := tx.Find(&categories)

	if result.Error != nil {
		slog.Warn("Error finding records from CATEGORY", "error", result.Error)
	}
	var resultCategories []map[string]interface{}
	for _, category := range categories {
		result := tx.Where("category_id = ? AND deleted_at IS NULL", category.ID).Find(&products)
		if result.Error != nil {
			slog.Warn("Error finding records from PRODUCT", "error", result.Error)
		}

		var resultProducts []map[string]interface{}
		for _, product := range products {
			resultProducts = append(resultProducts, map[string]interface{}{
				"description":  product.Description,
				"price":        product.Price,
				"product_name": product.ProductName,
				"uuid":         product.UUID,
			})
		}
		resultCategories = append(resultCategories, map[string]interface{}{
			"category_name": category.CategoryName,
			"products":      resultProducts,
		})
	}
	tx.Commit()
	return resultCategories
}

func (d *CategoryShopPostgres) Delete(categoryID string) error {
	var category models.Category
	tx := d.db.Begin()
	result := tx.Where("category_name = ?", categoryID).Delete(&category)
	if result.Error != nil {
		slog.Warn("Error in category's deleting method", "error", result.Error)
		return result.Error
	}
	tx.Commit()
	return result.Error
}

func (d *CategoryShopPostgres) Get(categoryID string) map[string]interface{} {
	var category models.Category
	var products []models.Product
	tx := d.db.Begin()
	result := tx.Where("category_name = ?", categoryID).First(&category)
	if result.Error != nil {
		slog.Warn("Error finding by CATEGORY_NAME", "error", result.Error)
	}

	res := tx.Where("category_id = ?", category.ID).First(&products)
	if res.Error != nil {
		slog.Warn("Error finding PRODUCTS which making relationships with CategoryID by CATEGORY_NAME", "error", res.Error)
	}

	var resultProducts []map[string]interface{}
	for _, product := range products {
		resultProducts = append(resultProducts, map[string]interface{}{
			"description":  product.Description,
			"price":        product.Price,
			"product_name": product.ProductName,
			"uuid":         product.UUID,
		})
	}
	resultCategory := map[string]interface{}{
		"category_name": category.CategoryName,
		"products":      resultProducts,
	}
	tx.Commit()
	return resultCategory
}

func (d *CategoryShopPostgres) Update(categoryID string, newCategory models.Category) error {
	var category models.Category
	tx := d.db.Begin()
	result := tx.Where("category_name = ? AND deleted_at IS NULL", categoryID).First(&category)
	if result.Error != nil {
		slog.Warn("Error in category's deleting method", "error", result.Error)
		return result.Error
	}
	category.CategoryName = newCategory.CategoryName
	category.UpdatedAt = time.Now()
	result = tx.Save(&category)
	tx.Commit()
	return result.Error
}
