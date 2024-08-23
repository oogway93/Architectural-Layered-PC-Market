package serviceShop

import (
	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/repository"
)

type CategoryShopService struct {
	repositoryShop repository.CategoryRepository
}

func NewServiceShopCategory(repo repository.CategoryRepository) *CategoryShopService {
	return &CategoryShopService{
		repositoryShop: repo,
	}
}

func (c *CategoryShopService) GetAll() ([]products.Category, error) {
	return nil, nil
}
