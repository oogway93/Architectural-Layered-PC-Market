package repositoryPostgresShop

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
	"github.com/shopspring/decimal"
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

func (d *ProductShopPostgres) GetByCategoryId(categoryID uint) string {
	tx := d.db.Begin()
	var existingCategory models.Category
	result := tx.Unscoped().Where("id = ?", categoryID).First(&existingCategory)
	if result.Error != nil {
		slog.Warn("Error finding records from CATEGORY", "error", result.Error)
	}
	tx.Commit()
	return existingCategory.CategoryName
}

func (d *ProductShopPostgres) Create(categoryID string, newProduct *models.Product) error {
	tx := d.db.Begin()

	var checkingProduct models.Product
	check := tx.Unscoped().Where("product_name = ?", newProduct.ProductName).First(&checkingProduct)
	if check.Error == nil {
		rawSQL := `UPDATE products SET deleted_at = NULL WHERE product_name = ?`
		tx.Exec(rawSQL, checkingProduct.ProductName)
		slog.Info("Restored PRODUCT", "productName", newProduct.ProductName)
	} else {
		if check.RowsAffected == 0 {
			var existingCategory models.Category
			result := tx.Where("category_name = ?", categoryID).First(&existingCategory)

			if result.Error != nil {
				slog.Warn("Error finding CATEGORY", "error", result.Error)
				tx.Rollback()
				return fmt.Errorf("failed to find CATEGORY")
			}

			newProduct.CategoryID = existingCategory.ID
			result = tx.Create(&newProduct)
			if result.Error != nil {
				slog.Warn("Error creating new PRODUCT: %v", "error", result.Error)
				tx.Rollback()
				return fmt.Errorf("failed to create PRODUCT")
			}

			slog.Info("Created new PRODUCT:"+"newProduct"+"in CATEGORY", newProduct.ProductName, existingCategory.CategoryName)
		}
	}

	tx.Commit()
	return nil
}

func (d *ProductShopPostgres) GetAll(categoryID string) ([]models.Product, []map[string]interface{}) {
	var products []models.Product
	var category models.Category
	tx := d.db.Begin()
	getCategoryID := tx.Where("category_name = ?", categoryID).First(&category)

	if getCategoryID.Error != nil {
		slog.Warn("Error finding records from CATEGORY", "error", getCategoryID.Error)
	}

	result := tx.Where("category_id = ?", category.ID).Find(&products)
	if result.Error != nil {
		slog.Warn("Error finding records from PRODUCT", "error", result.Error)
	}
	var resultProducts []map[string]interface{}
	for _, product := range products {
		resultProducts = append(resultProducts, map[string]interface{}{
			"description":   product.Description,
			"price":         product.Price,
			"product_name":  product.ProductName,
			"category_name": category.CategoryName,
			"uuid":          product.UUID,
		})
	}
	tx.Commit()
	return products, resultProducts
}

func (d *ProductShopPostgres) Delete(categoryID string, productID string) error {
	var product models.Product
	tx := d.db.Begin()
	result := tx.Where("product_name = ?", productID).Delete(&product)
	if result.Error != nil {
		slog.Warn("Error in products's deleting method", "error", result.Error)
		return result.Error
	}
	tx.Commit()
	return result.Error
}

func (d *ProductShopPostgres) Get(categoryID string, productID string) (models.Product, map[string]interface{}) {
	var product models.Product
	var category models.Category
	tx := d.db.Begin()
	getCategoryID := tx.Where("category_name = ?", categoryID).First(&category)

	if getCategoryID.Error != nil {
		slog.Warn("Error finding records from category", "error", getCategoryID.Error)
	}

	result := tx.Where("category_id = ? AND product_name = ?", category.ID, productID).First(&product)
	if result.Error != nil {
		slog.Warn("Error finding records from product", "error", result.Error)
	}
	resultProduct := map[string]interface{}{
		"uuid":          product.UUID,
		"category_name": category.CategoryName,
		"product_name":  product.ProductName,
		"price":         product.Price,
		"description":   product.Description,
	}
	tx.Commit()
	return product, resultProduct
}
func (d *ProductShopPostgres) Update(newCategoryName, productID string, newProduct models.Product) (map[string]interface{}, error) {
	var product models.Product
	var category models.Category
	tx := d.db.Begin()

	result := tx.Where("product_name = ?", productID).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	if newCategoryName != "" {
		getCategory := tx.Where("category_name = ?", newCategoryName).First(&category)

		if getCategory.Error != nil {
			slog.Warn("Error finding records from CATEGORY", "error", getCategory.Error)
		}
		product.CategoryID = category.ID
	}

	if newProduct.ProductName != "" {
		product.ProductName = newProduct.ProductName
	}
	if newProduct.Price.Cmp(decimal.NewFromInt(0)) > 0 {
		product.Price = newProduct.Price
	}
	if newProduct.Description != "" {
		product.Description = newProduct.Description
	}
	product.UpdatedAt = time.Now()
	result = tx.Save(&product)
	tx.Commit()
	resultProduct := map[string]interface{}{
		"uuid":          product.UUID,
		"categoryID":    product.CategoryID,
		"category_name": category.CategoryName,
		"product_name":  product.ProductName,
		"price":         product.Price,
		"description":   product.Description,
	}
	return resultProduct, result.Error
}
