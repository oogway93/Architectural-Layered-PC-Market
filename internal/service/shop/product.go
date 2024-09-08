package serviceShop

import (
	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/repository"
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
)

type ProductShopService struct {
	repo repository.ProductRepository
	cache                 repository.CacheRepository
}

func NewServiceShopProduct(repo repository.ProductRepository, cache repository.CacheRepository) *ProductShopService {
	return &ProductShopService{
		repo: repo,
		cache: cache,
	}
}

func (c *ProductShopService) Create(categoryID string, requestData *products.Product) {
	productModel := models.Product{
		ProductName: requestData.ProductName,
		Price:       requestData.Price,
		Description: requestData.Description,
	}
	c.repo.Create(categoryID, productModel)
}
func (c *ProductShopService) GetAll(categoryID string) []map[string]interface{} {
	result := c.repo.GetAll(categoryID)
	return result
}
func (c *ProductShopService) Delete(categoryID, productID string) error {
	result := c.repo.Delete(categoryID, productID)
	return result
}
func (c *ProductShopService) Get(categoryID, productID string) map[string]interface{} {
	result := c.repo.Get(categoryID, productID)
	return result
}
func (c *ProductShopService) Update(categoryID, productID string, requestData *products.Product) error {
	productModel := models.Product{
		ProductName: requestData.ProductName,
		Price:       requestData.Price,
		Description: requestData.Description,
	}
	newCategoryName := requestData.CategoryRel.CategoryName
	result := c.repo.Update(newCategoryName, productID, productModel)
	return result
}
