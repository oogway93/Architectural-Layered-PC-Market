package productsAPI

import (
	"github.com/shopspring/decimal"
)

type Product struct {
	ProductName string          `json:"product_name"`
	Price       decimal.Decimal `json:"price"`
	Description string          `json:"description"`
	CategoryRel Category        `json:"category_rel"`
}