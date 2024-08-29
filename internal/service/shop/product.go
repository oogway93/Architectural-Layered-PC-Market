package serviceShop

import (
	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/repository"
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
)

type ProductShopService struct {
	repositoryShopProduct repository.ProductRepository
}

func NewServiceShopProduct(repo repository.ProductRepository) *ProductShopService {
	return &ProductShopService{
		repositoryShopProduct: repo,
	}
}

func (c *ProductShopService) Create(categoryID string, requestData *products.Product) {
	productModel := models.Product{
		ProductName: requestData.ProductName,
		Price:       requestData.Price,
		Description: requestData.Description,
	}
	c.repositoryShopProduct.Create(categoryID, productModel)
}
func (c *ProductShopService) GetAll(categoryID string) []map[string]interface{} {
	result := c.repositoryShopProduct.GetAll(categoryID)
	return result
}
func (c *ProductShopService) Delete(categoryID, productID string) error {
	result := c.repositoryShopProduct.Delete(categoryID, productID)
	return result 
}
func (c *ProductShopService) Get(categoryID, productID string) string {
	result := c.repositoryShopProduct.Get(categoryID, productID)
	return result
}
func (c *ProductShopService) Update(categoryID, productID string, requestData *products.Product) error {
	productModel := models.Product{
		ProductName: requestData.ProductName,
		Price:       requestData.Price,
		Description: requestData.Description,
	}
	result := c.repositoryShopProduct.Update(categoryID, productID, productModel)
	return result
}
