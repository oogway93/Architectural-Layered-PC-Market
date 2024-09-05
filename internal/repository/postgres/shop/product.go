package repositoryPostgresShop

import (
	"fmt"
	"log"
	"time"

	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
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

func (d *ProductShopPostgres) Create(categoryID string, newProduct models.Product) error {
	tx := d.db.Begin()

	var checkingProduct models.Product
	check := tx.Unscoped().Where("product_name = ? AND deleted_at IS NOT NULL", newProduct.ProductName).First(&checkingProduct)
	if check.Error == nil {
		rawSQL := `UPDATE products SET deleted_at = NULL WHERE product_name = ?`
		tx.Exec(rawSQL, checkingProduct.ProductName)
		log.Printf("Restored product: %s", newProduct.ProductName)
	} else {
		var existingCategory models.Category
		result := tx.Where("category_name = ? AND deleted_at IS NULL", categoryID).First(&existingCategory)

		if result.Error != nil {
			log.Printf("Error finding category: %v", result.Error)
			tx.Rollback()
			return fmt.Errorf("failed to find category")
		}

		newProduct.CategoryID = existingCategory.ID
		result = tx.Create(&newProduct)
		if result.Error != nil {
			log.Printf("Error creating new product: %v", result.Error)
			tx.Rollback()
			return fmt.Errorf("failed to create product")
		}

		log.Printf("Created new product: %s in category: %s", newProduct.ProductName, existingCategory.CategoryName)
	}

	tx.Commit()
	return nil
}

func (d *ProductShopPostgres) GetAll(categoryID string) []map[string]interface{} {
	var products []models.Product
	var category models.Category
	tx := d.db.Begin()
	getCategoryID := tx.Where("category_name = ?", categoryID).First(&category)

	if getCategoryID.Error != nil {
		log.Printf("Error finding records from category: %v", getCategoryID.Error)
	}

	result := tx.Where("category_id = ?", category.ID).Find(&products)
	if result.Error != nil {
		log.Printf("Error finding records from product: %v", result.Error)
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
	return resultProducts
}

func (d *ProductShopPostgres) Delete(categoryID string, productID string) error {
	var product models.Product
	// var category models.Category
	tx := d.db.Begin()

	// getCategoryID := tx.Where("category_name = ? A ND deleted_at IS NULL", categoryID).First(&category)

	// if getCategoryID.Error != nil {
	// 	log.Printf("Error finding records from category: %v", getCategoryID.Error)
	// }

	result := tx.Where("product_name = ?", productID).Delete(&product)
	if result.Error != nil {
		return result.Error
	}
	tx.Commit()
	return result.Error
}

func (d *ProductShopPostgres) Get(categoryID string, productID string) map[string]interface{} {
	var product models.Product
	var category models.Category
	tx := d.db.Begin()
	getCategoryID := tx.Where("category_name = ? AND deleted_at IS NULL", categoryID).First(&category)

	if getCategoryID.Error != nil {
		log.Printf("Error finding records from category: %v", getCategoryID.Error)
	}

	result := tx.Where("category_id = ? AND product_name = ? AND deleted_at IS NULL", category.ID, productID).First(&product)
	if result.Error != nil {
		log.Printf("Error finding records from product: %v", result.Error)
	}
	resultProduct := map[string]interface{}{
		"uuid":          product.UUID,
		"category_name": category.CategoryName,
		"product_name":  product.ProductName,
		"price":         product.Price,
		"description":   product.Description,
	}
	tx.Commit()
	return resultProduct
}
func (d *ProductShopPostgres) Update(newCategoryName, productID string, newProduct models.Product) error {
	var product models.Product
	var category models.Category
	tx := d.db.Begin()
	
	result := tx.Where("product_name = ? AND deleted_at IS NULL", productID).First(&product)
	if result.Error != nil {
		return result.Error
	}

	if newCategoryName != "" {
		getCategory := tx.Where("category_name = ? AND deleted_at IS NULL", newCategoryName).First(&category)
		
		if getCategory.Error != nil {
			log.Printf("Error finding records from category: %v", getCategory.Error)
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
	return result.Error
}
