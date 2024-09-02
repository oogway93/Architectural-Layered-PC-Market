package serviceShop

import (
	"log"

	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/repository"
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
	"github.com/shopspring/decimal"
)

type OrderShopService struct {
	repositoryShopOrder repository.OrderRepository
}

func NewServiceShopOrder(repo repository.OrderRepository) *OrderShopService {
	return &OrderShopService{
		repositoryShopOrder: repo,
	}
}

func (c *OrderShopService) Create(userID string, requestData *products.Order) {
	deliveryModel := models.Delivery{
		FullName:      requestData.DeliveryRel.FullName,
		Postcode:      requestData.DeliveryRel.Postcode,
		Country:       requestData.DeliveryRel.Country,
		City:          requestData.DeliveryRel.City,
		DeliveryPrice: requestData.DeliveryRel.DeliveryPrice,
	}

	var orderItems []*models.OrderItem
	for _, productItem := range requestData.OrderItemsRel {
		resultProduct := c.repositoryShopOrder.FetchProductID(productItem.ProductRel.ProductName)
		unitPrice, ok := resultProduct["price"].(decimal.Decimal)
		if ok {
			orderItem := models.OrderItem{
				ProductID: resultProduct["id"].(uint),
				Quantity:  productItem.Quantity,
				UnitPrice: unitPrice,
			}
			orderItems = append(orderItems, &orderItem)
		}
	}

	// orderModel.OrderItems = orderItems

	c.repositoryShopOrder.CreateDelivery(&deliveryModel)

	deliveryID, err := c.repositoryShopOrder.LastRow()
	if err != nil {
		log.Fatalf("Error in getting last row from delivery: %v", err.Error())
	}

	c.repositoryShopOrder.CreateOrderAndOrderItems(userID, deliveryID, orderItems)

}

func (s *OrderShopService) GetAll() []map[string]interface{}                       { return nil }
func (s *OrderShopService) Get(orderID string) map[string]interface{}              { return nil }
func (s *OrderShopService) Update(orderID string, requestData *models.Order) error { return nil }
func (s *OrderShopService) Delete(orderID string) error                            { return nil }
