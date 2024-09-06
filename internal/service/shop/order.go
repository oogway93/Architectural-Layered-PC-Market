package serviceShop

import (
	"log"
	"time"

	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/repository"
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
	"github.com/shopspring/decimal"
)

const (
	Shipped   = "Shipped"
	Delivered = "Delivered"
	PickedUp  = "Picked_up"
	Completed = "Completed"
)

type OrderShopService struct {
	repositoryShopOrder repository.OrderRepository
}

func NewServiceShopOrder(repo repository.OrderRepository) *OrderShopService {
	return &OrderShopService{
		repositoryShopOrder: repo,
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
		resultProduct := s.repositoryShopOrder.FetchProductID(productItem.ProductRel.ProductName)
		unitPrice, ok := resultProduct["price"].(decimal.Decimal)
		if !ok {
			log.Printf("Failed to convert price to decimal for product %s", productItem.ProductRel.ProductName)
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
		log.Println("No order items found")
		return
	}

	s.repositoryShopOrder.CreateDelivery(&deliveryModel)

	deliveryID, err := s.repositoryShopOrder.LastRow()
	if err != nil {
		log.Fatalf("Error in getting last row from delivery: %v", err.Error())
	}

	order := s.repositoryShopOrder.CreateOrderAndOrderItems(userID, deliveryID, orderItems)
	go s.autoUpdateStatus(order.ID)
}

func (s *OrderShopService) autoUpdateStatus(orderID uint) {
	statusUpdates := []struct {
		status string
		delay  time.Duration
	}{
		{Delivered, 10 * time.Minute},
		{Shipped, 30 * time.Minute},
		{PickedUp, 50 * time.Minute},
		{Completed, 50 * time.Minute + 1 * time.Second},
	}
	for _, update := range statusUpdates {
		time.Sleep(update.delay)
		s.repositoryShopOrder.UpdateOrderStatus(orderID, update.status)
		time.Sleep(1 * time.Second)
		if update.status == "Completed" {
			s.repositoryShopOrder.Delete(orderID)
		}
	}
}
func (s *OrderShopService) GetAll(userID string) []map[string]interface{} {
	result := s.repositoryShopOrder.GetAll(userID)
	return result
}
func (s *OrderShopService) Get(orderID string) map[string]interface{}              { return nil }
func (s *OrderShopService) Update(orderID string, requestData *models.Order) error { return nil }
func (s *OrderShopService) Delete(orderID string) error                            { return nil }
