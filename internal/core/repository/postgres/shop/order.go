package repositoryPostgresShop

import (
	"fmt"
	"log/slog"

	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
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

func (d *OrderShopPostgres) CreateOrderAndOrderItems(userID string, deliveryID uint, newItems []*models.OrderItem) *models.Order {
	var user models.User
	tx := d.db.Begin()
	result := tx.Where("login = ?", userID).First(&user)
	if result.Error != nil {
		slog.Warn("Error finding LOGIN from user", "error", result.Error)
	}

	newOrder := models.Order{
		UserID:     user.ID,
		DeliveryID: deliveryID,
		Status:     "In_process",
	}

	var total decimal.Decimal
	for _, item := range newItems {
		quantityInt64 := int64(item.Quantity)
		total = total.Add(item.UnitPrice.Mul(decimal.NewFromInt(quantityInt64)))
	}

	newOrder.Total = total

	result = tx.Create(&newOrder)
	if result.Error != nil {
		slog.Warn("Error creating new CATEGORY", "error", result.Error)
		tx.Rollback()
	} else {
		slog.Info("Created new order")
	}

	for _, item := range newItems {
		item.OrderID = newOrder.ID
		result := tx.Create(&item)
		if result.Error != nil {
			slog.Warn("Error creating ORDER ITEMS", "error", result.Error)
			tx.Rollback()
		} else {
			slog.Info("Created new ORDER ITEMS")
		}
	}
	tx.Commit()
	return &newOrder
}
func (d *OrderShopPostgres) UpdateOrderStatus(orderID string, newStatus string) {
	var order models.Order
	result := d.db.Where("uuid = ?", orderID).First(&order)
	if result.RowsAffected == 0 {
		slog.Warn("ORDER not found")
	}

	order.Status = newStatus

	if err := d.db.Save(&order).Error; err != nil {
		slog.Warn("Failed to save ORDER", "error", err)
	}
}

func (d *OrderShopPostgres) CreateDelivery(newDelivery *models.Delivery) {
	tx := d.db.Begin()

	result := tx.Create(&newDelivery)
	if result.Error != nil {
		slog.Warn("Error creating new ORDER-DELIVERY", "error", result.Error)
	} else {
		slog.Info("Created new ORDER-DELIVERY")
	}

	tx.Commit()
}

func (d *OrderShopPostgres) LastRow() (uint, error) {
	var delivery models.Delivery
	tx := d.db.Begin()

	result := tx.Limit(1).Order("id desc").First(&delivery)
	if result.Error != nil {
		return 0, fmt.Errorf("failed to get last inserted order: %v", result.Error)
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
		slog.Warn("Error finding LOGIN from USER", "error", result.Error)
	}
	result = tx.Where("user_id = ?", user.ID).Find(&orders)
	if result.Error != nil {
		slog.Warn("Error finding records from ORDER", "error", result.Error)
	}

	slog.Info("Found ORDERS", "orderLength", len(orders))

	var resultOrders []map[string]interface{}
	for _, order := range orders {
		slog.Info("Processing order", "orderId", order.ID)

		resultOrder := make(map[string]interface{})
		resultOrder["status"] = order.Status
		resultOrder["total_price"] = order.Total

		result = tx.Preload("Order").Preload("Product").Where("order_id = ?", order.ID).Find(&orderItems)
		if result.Error != nil {
			slog.Warn("Error finding records from order items", "error", result.Error)
		}

		var resultOrderItems []map[string]interface{}
		for _, item := range orderItems {
			slog.Info("Processing item", "itemID", item.ID)

			result = tx.Preload("Category").Where("id = ?", item.Product.ID).Find(&product)
			if result.Error != nil {
				slog.Warn("Error finding records from order items", "error", result.Error)
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

		slog.Info("Added ORDER to resultOrders")
	}

	tx.Commit()
	return resultOrders
}
func (d *OrderShopPostgres) Get()    {}
func (d *OrderShopPostgres) Update() {}
func (d *OrderShopPostgres) Delete(orderID string) error {
	var order models.Order
	tx := d.db.Begin()
	result := tx.Where("uuid = ?", orderID).Delete(&order)
	if result.Error != nil {
		slog.Warn("Error in DELETE method ORDER", "error", result.Error)
	}
	tx.Commit()
	return result.Error
}
func (d *OrderShopPostgres) FetchProductID(productName string) map[string]interface{} {
	var product models.Product
	tx := d.db.Begin()

	result := tx.Where("product_name = ?", productName).First(&product)
	if result.Error != nil {
		slog.Warn("Error finding PRODUCT", "error", result.Error)
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
