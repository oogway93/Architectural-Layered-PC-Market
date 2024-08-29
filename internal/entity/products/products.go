package products

import (
	"github.com/shopspring/decimal"
)

type Product struct {
	// UUID        uuid.UUID       `json:"uuid"`
	ProductName string          `json:"product_name"`
	Price       decimal.Decimal `json:"price"`
	Description string          `json:"description"`
	// CategoryRel Category        `json:"category_name"`
}
