package service

import (
	// "github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/entity/user"
	"github.com/oogway93/golangArchitecture/internal/repository"
	"github.com/oogway93/golangArchitecture/internal/service/shop"
	serviceUser "github.com/oogway93/golangArchitecture/internal/service/user"
)

type ServiceCategory interface {
	Create(requestData *products.Category) ()
}

type ServiceUser interface {
	Create(requestData *user.User) ()
	GetAll() []map[string]interface{}
	Update(loginID string, requestData *user.UserUpdated) (error)
	Delete(loginID string) (error)
}

type Service struct {
	ServiceCategory ServiceCategory
	ServiceUser ServiceUser
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		ServiceCategory: serviceShop.NewServiceShopCategory(repo.CategoryRepository),
		ServiceUser: serviceUser.NewServiceUser(repo.UserRepository),
	}
}
