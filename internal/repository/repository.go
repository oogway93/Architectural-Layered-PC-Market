package repository

import (
	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
)


type CategoryRepository interface {
	GetAll() ([]products.Category,error)
}

type UserRepository interface {
	Create(newUser models.User) ()
	GetAll() []map[string]interface{}
	Update(loginId string, newUser models.User) (error)
}

type Repository struct {
	CategoryRepository CategoryRepository
	UserRepository UserRepository
}


