package repositoryPostgresShop

import (
	"fmt"
	"log"

	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
	"github.com/shopspring/decimal"
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

	var total decimal.Decimal
	for _, item := range newItems {
		quantityInt64 := int64(item.Quantity)
		total = total.Add(item.UnitPrice.Mul(decimal.NewFromInt(quantityInt64)))
		// result := tx.Create(&item)
	}

	newOrder.Total = total

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

func (d *OrderShopPostgres) GetAll(userID string) []map[string]interface{} {
	var orders []models.Order
	var orderItems []models.OrderItem
	var user models.User
	var product models.Product
	tx := d.db.Begin()
	result := tx.Where("login = ?", userID).First(&user)
	if result.Error != nil {
		log.Printf("Error finding LOGIN from user: %v", result.Error)
	}
	result = tx.Where("user_id = ?", user.ID).Find(&orders)
	if result.Error != nil {
		log.Printf("Error finding records from order: %v", result.Error)
	}

	log.Printf("Found %d orders", len(orders))

	var resultOrders []map[string]interface{}
	for _, order := range orders {
		log.Printf("Processing order: %v", order.ID)

		resultOrder := make(map[string]interface{})
		resultOrder["status"] = order.Status
		resultOrder["total_price"] = order.Total

		result = tx.Preload("Order").Preload("Product").Where("order_id = ?", order.ID).Find(&orderItems)
		if result.Error != nil {
			log.Printf("Error finding records from order items: %v", result.Error)
		}

		var resultOrderItems []map[string]interface{}
		for _, item := range orderItems {
			log.Printf("Processing item: %v", item.ID)

			result = tx.Preload("Category").Where("id = ?", item.Product.ID).Find(&product)
			if result.Error != nil {
				log.Printf("Error finding records from order items: %v", result.Error)
			}

			resultOrderItems = append(resultOrderItems, map[string]interface{}{
				"category":    product.Category.CategoryName,
				"product":     item.Product.ProductName,
				"description": item.Product.Description,
				"quantity":    item.Quantity,
				"unit_price":  item.UnitPrice,
				"uuid":        item.Product.UUID,
			})
		}

		resultOrder["order_items"] = resultOrderItems
		resultOrders = append(resultOrders, resultOrder)

		log.Printf("Added order to resultOrders")
	}

	tx.Commit()
	return resultOrders
}
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
