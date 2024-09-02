package repositoryPostgresShop

import (
	"fmt"
	"log"

	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
	"gorm.io/gorm"
)

type OrderShopPostgres struct {
	db *gorm.DB
}

func NewRepositoryOrderShop(db *gorm.DB) *OrderShopPostgres {
	return &OrderShopPostgres{
		db: db,
	}
}

func (d *OrderShopPostgres) CreateOrderAndOrderItems(userID string, deliveryID uint, newItems []*models.OrderItem) {
	tx := d.db.Begin()
	var user models.User
	result := tx.Where("login = ?", userID).First(&user)
	if result.Error != nil {
		log.Printf("Error finding LOGIN from user: %v", result.Error)
	}

	newOrder := models.Order{
		UserID:     user.ID,
		DeliveryID: deliveryID,
		Status:     "in_process",
	}

	result = tx.Create(&newOrder)
	if result.Error != nil {
		log.Printf("Error creating new category: %v", result.Error)
		tx.Rollback()
	} else {
		log.Printf("Created new order")
	}

	for _, item := range newItems {
		item.OrderID = newOrder.ID
		result := tx.Create(&item)
		if result.Error != nil {
			log.Printf("Error creating order item: %v", result.Error)
			tx.Rollback()
		} else {
			log.Printf("Created new order items")
		}
	}

	tx.Commit()
}

// FIXME: создает две записи(видимо сколько в список добавляешь предметов для заказа, столько он и сохраняет одинаковых записей в бд о доставке)
func (d *OrderShopPostgres) CreateDelivery(newDelivery *models.Delivery) {
	tx := d.db.Begin()

	result := tx.Create(&newDelivery)
	if result.Error != nil {
		log.Printf("Error creating new order-delivery: %v", result.Error)
	} else {
		log.Printf("Created new order-delivery")
	}

	tx.Commit()
}

func (d *OrderShopPostgres) LastRow() (uint, error) {
	var delivery models.Delivery
	tx := d.db.Begin()

	result := tx.Limit(1).Order("id desc").First(&delivery)
	if result.Error != nil {
		return 0, fmt.Errorf("failed to get last inserted order: %w", result.Error)
	}

	tx.Commit()
	return delivery.ID, nil
}

func (d *OrderShopPostgres) GetAll() {}
func (d *OrderShopPostgres) Get()    {}
func (d *OrderShopPostgres) Update() {}
func (d *OrderShopPostgres) Delete() {}
func (d *OrderShopPostgres) FetchProductID(productName string) map[string]interface{} {
	var product models.Product
	tx := d.db.Begin()

	result := tx.Where("product_name = ?", productName).First(&product)
	if result.Error != nil {
		log.Printf("Error finding product: %v", result.Error)
		tx.Rollback()
	}

	resultProduct := map[string]interface{}{
		"id":            product.ID,
		"category_name": product.Category.CategoryName,
		"product_name":  product.ProductName,
		"price":         product.Price,
	}
	tx.Commit()

	return resultProduct
}
