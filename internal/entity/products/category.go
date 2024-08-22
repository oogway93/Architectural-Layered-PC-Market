package products

type Category struct {
	Name string `json:"name", db:"name"`
	Description string `json:"description", db:"description"`
}
