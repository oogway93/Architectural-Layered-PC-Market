package products

type ProductList struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Price int64 `json:"price"`
	Products []
}
