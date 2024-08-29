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

    // Check if category was found
    if existingCategory.ID == 0 {
        // If category doesn't exist, create it
        newCategory := models.Category{
            ID:         0, // Set to 0 to allow auto-increment
            CategoryName:       categoryID,
        }
        result = d.db.Create(&newCategory)
        if result.Error != nil {
            log.Printf("Error creating new category: %v", result.Error)
            tx.Rollback()
            return fmt.Errorf("failed to create category")
        }
        existingCategory = newCategory
    }

    // Associate product with category
    newProduct.CategoryID = existingCategory.ID

    // Create new product
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

func (d *ProductShopPostgres) GetAll(categoryID string) []map[string]interface{} {return nil}
func (d *ProductShopPostgres) Delete(categoryID string, productID string) error {return nil}
func (d *ProductShopPostgres) Get(categoryID string, productID string) string {return "nil"}
func (d *ProductShopPostgres) Update(categoryID string, productID string, newProduct models.Product) error {return nil}
