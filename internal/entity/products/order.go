package products

import "github.com/shopspring/decimal"

type Order struct {
	DeliveryRel   Delivery     `json:"delivery_rel"`
	OrderItemsRel []OrderItems `json:"order_items_rel"`
}

type Delivery struct {
	FullName      string          `json:"full_name"`
	Postcode      string          `json:"postcode"`
	Country       string          `json:"country"`
	City          string          `json:"city"`
	DeliveryPrice decimal.Decimal `json:"delivery_price"`
}

type OrderItems struct {
	Quantity   int          `json:"quantity"`
	ProductRel ProductItems `json:"product_rel"`
}

type ProductItems struct {
	ProductName string `json:"product_name"`
}
