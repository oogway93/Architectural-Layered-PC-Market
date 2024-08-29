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

	var existingCategory models.Category
	result := d.db.Where("category_name = ?", categoryID).First(&existingCategory)

	if result.Error != nil {
		log.Printf("Error finding category: %v", result.Error)
		tx.Rollback()
		return fmt.Errorf("failed to find category")
	}

	if existingCategory.ID == 0 {
		newCategory := models.Category{
			ID:           0,
			CategoryName: categoryID,
		}
		result = d.db.Create(&newCategory)
		if result.Error != nil {
			log.Printf("Error creating new category: %v", result.Error)
			tx.Rollback()
			return fmt.Errorf("failed to create category")
		}
		existingCategory = newCategory
	}

	newProduct.CategoryID = existingCategory.ID

	result = tx.Create(&newProduct)
	if result.Error != nil {
		log.Printf("Error creating new product: %v", result.Error)
		tx.Rollback()
		return fmt.Errorf("failed to create product")
	}

	log.Printf("Created new product: %s in category: %s", newProduct.ProductName, existingCategory.CategoryName)

	tx.Commit()
	return nil
}

func (d *ProductShopPostgres) GetAll(categoryID string) []map[string]interface{} {
	var products []models.Product
	var category models.Category
	tx := d.db.Begin()
	getCategoryID := d.db.Where("category_name = ? AND deleted_at IS NULL", categoryID).First(&category)

	if getCategoryID.Error != nil {
		log.Printf("Error finding records from category: %v", getCategoryID.Error)
	}

	result := d.db.Where("category_id = ? AND deleted_at IS NULL", category.ID).Find(&products)
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
func (d *ProductShopPostgres) Delete(categoryID string, productID string) error { return nil }
func (d *ProductShopPostgres) Get(categoryID string, productID string) string   { return "nil" }
func (d *ProductShopPostgres) Update(categoryID string, productID string, newProduct models.Product) error {
	return nil
}
