package products

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	UUID        uuid.UUID
	ProductName string          `json:"product_name"`
	Price       decimal.Decimal `json:"price"`
	Description string          `json:"description"`
	CategoryRel Category        `json:"category_rel"`
}
