package serviceShop

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/oogway93/golangArchitecture/internal/core/entity/products"
	"github.com/oogway93/golangArchitecture/internal/core/repository"
	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
	"github.com/oogway93/golangArchitecture/internal/core/utils"
	"github.com/shopspring/decimal"
)

const (
	Shipped   = "Shipped"
	Delivered = "Delivered"
	PickedUp  = "Picked_up"
	Completed = "Completed"
)

type OrderShopService struct {
	repo  repository.OrderRepository
	cache repository.CacheRepository
}

func NewServiceShopOrder(repo repository.OrderRepository, cache repository.CacheRepository) *OrderShopService {
	return &OrderShopService{
		repo:  repo,
		cache: cache,
	}
}

func (s *OrderShopService) Create(userID string, requestData *products.Order) {
	deliveryModel := models.Delivery{
		FullName:      requestData.DeliveryRel.FullName,
		Postcode:      requestData.DeliveryRel.Postcode,
		Country:       requestData.DeliveryRel.Country,
		City:          requestData.DeliveryRel.City,
		DeliveryPrice: requestData.DeliveryRel.DeliveryPrice,
	}

	var orderItems []*models.OrderItem
	for _, productItem := range requestData.OrderItemsRel {
		resultProduct := s.repo.FetchProductID(productItem.ProductRel.ProductName)
		unitPrice, ok := resultProduct["price"].(decimal.Decimal)
		if !ok {
			slog.Info("Failed to convert price to decimal for PRODUCT", "productName", productItem.ProductRel.ProductName)
			continue
		}
		orderItem := models.OrderItem{
			ProductID: resultProduct["id"].(uint),
			Quantity:  productItem.Quantity,
			UnitPrice: unitPrice,
		}
		orderItems = append(orderItems, &orderItem)

	}
	if len(orderItems) == 0 {
		slog.Info("No order items found")
		return
	}

	s.repo.CreateDelivery(&deliveryModel)

	deliveryID, err := s.repo.LastRow()
	if err != nil {
		slog.Warn("Error in getting last row from delivery", "error", err.Error())
	}

	order := s.repo.CreateOrderAndOrderItems(userID, deliveryID, orderItems)
	go s.autoUpdateStatus(order.UUID.String())
}

func (s *OrderShopService) autoUpdateStatus(orderID string) {
	statusUpdates := []struct {
		status string
		delay  time.Duration
	}{
		{Delivered, 10 * time.Minute},
		{Shipped, 30 * time.Minute},
		{PickedUp, 50 * time.Minute},
		{Completed, 50*time.Minute + 1*time.Second},
	}

	for _, update := range statusUpdates {
		time.Sleep(update.delay)
		s.repo.UpdateOrderStatus(orderID, update.status)
		time.Sleep(1 * time.Second)
		if update.status == "Completed" {
			s.repo.Delete(orderID)
		}
	}
}
func (s *OrderShopService) GetAll(userID string) []map[string]interface{} {
	var orders []map[string]interface{}
	key := fmt.Sprintf("user:%s::orders", userID)
	resultOrders, err := s.cache.Get(key)
	if err == nil {
		err := utils.Deserialize(resultOrders, &orders)
		if err != nil {
			return nil
		}
		return orders
	}
	orders = s.repo.GetAll(userID)
	if orders != nil {
		ordersSerialized, err := utils.Serialize(orders)
		if err != nil {
			slog.Warn("serialization incorrect")
		}
		err = s.cache.Set(key, ordersSerialized, ttl)
		if err != nil {
			slog.Warn("set cache incorrect")
		}

		return orders
	}
	return nil
}
func (s *OrderShopService) remakeOrders(userID string) {
	var orders []map[string]interface{}
	key := fmt.Sprintf("user:%s::orders", userID)
	orders = s.repo.GetAll(userID)
	if orders != nil {
		ordersSerialized, err := utils.Serialize(orders)
		if err != nil {
			slog.Warn("serialization incorrect")
		}
		err = s.cache.Set(key, ordersSerialized, ttl)
		if err != nil {
			slog.Warn("set cache incorrect")
		}
	}
}
func (s *OrderShopService) Get(orderID string) map[string]interface{}              { return nil }
func (s *OrderShopService) Update(orderID string, requestData *models.Order) error { return nil }
func (s *OrderShopService) Delete(userID, orderID string) error {
	err := s.repo.Delete(orderID)
	if err != nil {
		return fmt.Errorf("error in Delete  method category repo postgres")
	}

	key := fmt.Sprintf("user:%s::orders", userID)
	err = s.cache.Delete(key)
	if err != nil {
		slog.Warn("error in Delete method order cache, because haven't find a key from the redis storage", "error", err.Error)
	}

	s.remakeOrders(userID)
	return err
}
