package serviceShop

import (
	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/repository"
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
)

type CategoryShopService struct {
	repositoryShopCategory repository.CategoryRepository
}

func NewServiceShopCategory(repo repository.CategoryRepository) *CategoryShopService {
	return &CategoryShopService{
		repositoryShopCategory: repo,
	}
}

func (c *CategoryShopService) Create(requestData *products.Category) {
	categoryModel := models.Category{
		CategoryName: requestData.CategoryName,
	}
	c.repositoryShopCategory.Create(categoryModel)
}

func (c *CategoryShopService) GetAll() []map[string]interface{} {
	result := c.repositoryShopCategory.GetAll()
	return result
}
func (c *CategoryShopService) Get(categoryID string) string {
	result := c.repositoryShopCategory.Get(categoryID)
	return result
}

func (c *CategoryShopService) Delete(categoryID string) error {
	result := c.repositoryShopCategory.Delete(categoryID)
	return result
}
func (c *CategoryShopService) Update(categoryID string, requestData products.Category) error {
	categoryModel := models.Category{
		CategoryName: requestData.CategoryName,
	}
	result := c.repositoryShopCategory.Update(categoryID, categoryModel)
	return result
}
