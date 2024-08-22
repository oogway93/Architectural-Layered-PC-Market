package serviceShop

import (
	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/repository"
)

type CategoryShopService struct {
	repo repository.CategoryRepository
}

func NewCategoryShopService(repo repository.CategoryRepository) *CategoryShopService {
	return &CategoryShopService{
		repo: repo,
	}
}

func GetAll(s *CategoryShopService) ([]products.Category, error) {
	return nil, nil
}
