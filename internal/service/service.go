package service

import (
	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/entity/user"
	"github.com/oogway93/golangArchitecture/internal/repository"
	serviceAuth "github.com/oogway93/golangArchitecture/internal/service/auth"
	"github.com/oogway93/golangArchitecture/internal/service/shop"
	"github.com/oogway93/golangArchitecture/internal/service/user"
)

type ServiceCategory interface {
	Create(requestData *products.Category)
	GetAll() []map[string]interface{}
	Get(categoryID string) map[string]interface{}
	Update(categoryID string, requestData *products.Category) error
	Delete(categoryID string) error
}

type ServiceProduct interface {
	Create(categoryID string, requestData *products.Product)
	GetAll(categoryID string) []map[string]interface{}
	Get(categoryID string, productID string) map[string]interface{}
	Update(categoryID string, productID string, requestData *products.Product) error
	Delete(categoryID, productID string) error
}
type ServiceUser interface {
	Create(requestData *user.User)
	GetAll() []map[string]interface{}
	Get(login string) map[string]interface{}
	Update(loginID string, requestData *user.UserUpdated) error
	Delete(loginID string) error
}
type ServiceAuth interface {
	Login(requestData *user.AuthInput) bool
}
type Service struct {
	ServiceCategory ServiceCategory
	ServiceProduct  ServiceProduct
	ServiceUser     ServiceUser
	ServiceAuth     ServiceAuth
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		ServiceCategory: serviceShop.NewServiceShopCategory(repo.CategoryRepository),
		ServiceProduct:  serviceShop.NewServiceShopProduct(repo.ProductRepository),
		ServiceUser:     serviceUser.NewServiceUser(repo.UserRepository),
		ServiceAuth:     serviceAuth.NewServiceAuth(repo.AuthRepository),
	}
}
