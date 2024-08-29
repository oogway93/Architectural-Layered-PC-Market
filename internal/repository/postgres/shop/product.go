package repositoryPostgresShop

import (
	"fmt"
	"log"

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

		// if existingCategory.ID == 0 {
		// 	newCategory := models.Category{
		// 		ID:           0,
		// 		CategoryName: categoryID,
		// 	}
		// 	result = tx.Create(&newCategory)
		// 	if result.Error != nil {
		// 		log.Printf("Error creating new category: %v", result.Error)
		// 		tx.Rollback()
		// 		return fmt.Errorf("failed to create category")
		// 	}
		// 	existingCategory = newCategory
		// }

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
	getCategoryID := tx.Where("category_name = ? AND deleted_at IS NULL", categoryID).First(&category)

	if getCategoryID.Error != nil {
		log.Printf("Error finding records from category: %v", getCategoryID.Error)
	}

	result := tx.Where("category_id = ? AND deleted_at IS NULL", category.ID).Find(&products)
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

	// getCategoryID := tx.Where("category_name = ? AND deleted_at IS NULL", categoryID).First(&category)

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
func (d *ProductShopPostgres) Update(categoryID string, productID string, newProduct models.Product) error {
	return nil
}
