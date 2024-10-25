package service

import (
	"github.com/oogway93/golangArchitecture/internal/core/entity/products"
	"github.com/oogway93/golangArchitecture/internal/core/entity/user"

	// "github.com/oogway93/golangArchitecture/internal/repository"
	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
	// "github.com/oogway93/golangArchitecture/internal/service/auth"
	// "github.com/oogway93/golangArchitecture/internal/service/shop"
	// "github.com/oogway93/golangArchitecture/internal/service/user"
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
	GetAll(reqFrom string, categoryID string) ([]models.Product, []map[string]interface{})
	Get(categoryID string, productID string) map[string]interface{}
	Update(newCategoryName, productID string, requestData *products.Product)  error
	Delete(categoryID, productID string) error
}

type ServiceOrder interface {
	Create(userID string, requestData *products.Order)
	GetAll(userID string) []map[string]interface{}
	Get(orderID string) map[string]interface{}
	Update(orderID string, requestData *models.Order) error
	Delete(userID, orderID string) error
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
	ServiceProduct  ServiceProduct
	ServiceCategory ServiceCategory
	ServiceOrder    ServiceOrder
	ServiceUser     ServiceUser
	ServiceAuth     ServiceAuth
}

// func NewService(repo *repository.Repository, cache *repository.Cache) *Service {
// 	return &Service{
// 		ServiceProduct:  serviceShop.NewServiceShopProduct(repo.ProductRepository),
// 		ServiceCategory: serviceShop.NewServiceShopCategory(repo.CategoryRepository, cache.CacheRepository),
// 		ServiceOrder:    serviceShop.NewServiceShopOrder(repo.OrderRepository),
// 		ServiceUser:     serviceUser.NewServiceUser(repo.UserRepository),
// 		ServiceAuth:     serviceAuth.NewServiceAuth(repo.AuthRepository),
// 	}
// }
