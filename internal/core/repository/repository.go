package repository

import (
	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
)

type CategoryRepository interface {
	Create(newCategory *models.Category)
	GetAll() ([]models.Category, []map[string]interface{})
	Delete(categoryID string) error
	Get(categoryID string) (models.Category, map[string]interface{})
	Update(categoryID string, newCategory models.Category) error
}

type ProductRepository interface {
	GetByCategoryId(categoryID uint) string
	Create(categoryID string, newProduct *models.Product) error
	GetAll(categoryID string) ([]models.Product, []map[string]interface{})
	Delete(categoryID string, productID string) error
	Get(categoryID string, productID string) map[string]interface{}
	Update(newCategoryName, productID string, newProduct models.Product) (map[string]interface{}, error)
}

type OrderRepository interface {
	CreateOrderAndOrderItems(userID string, deliveryID uint, newItems []*models.OrderItem) *models.Order
	CreateDelivery(newDelivery *models.Delivery)
	GetAll(userID string) []map[string]interface{}
	Get()
	Update()
	Delete(orderID string) error
	FetchProductID(productName string) map[string]interface{}
	LastRow() (uint, error)
	UpdateOrderStatus(orderID string, newStatus string)
}

type UserRepository interface {
	Create(newUser models.User)
	GetAll() []map[string]interface{}
	Get(loginID string) map[string]interface{}
	Update(loginID string, newUser models.User) error
	Delete(loginID string) error
}

type AuthRepository interface {
	Login(loginID string) map[string]interface{}
}

type CacheRepository interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
	DeleteByPrefix(prefix string) error
	Close() error
}

type Repository struct {
	ProductRepository  ProductRepository
	CategoryRepository CategoryRepository
	OrderRepository    OrderRepository
	UserRepository     UserRepository
	AuthRepository     AuthRepository
}
type Cache struct {
	CacheRepository CacheRepository
}
