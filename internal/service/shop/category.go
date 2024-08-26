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
