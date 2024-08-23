package service

import (
	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/repository"
	"github.com/oogway93/golangArchitecture/internal/service/shop"
)

type ServiceCategory interface {
	GetAll() ([]products.Category, error)
}

type Service struct {
	ServiceCategory ServiceCategory
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		ServiceCategory: serviceShop.NewServiceShopCategory(repo.CategoryRepository),
	}
}
