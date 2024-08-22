package service

import (
	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/repository"
	"github.com/oogway93/golangArchitecture/internal/service/shop/serviceShop"
)

type CategoryService interface {
	GetAll() ([]products.Category, error)
}

type Service struct {
	CategoryService
}

func NewService(repo *repository.CategoryRepository) *Service {
	return &Service{
		CategoryService: serviceShop.NewCategoryShopService(repo),
	}
}
