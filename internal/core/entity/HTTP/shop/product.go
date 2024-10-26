package productsHTTP

import "github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"

type Product struct {
	Product models.Product
	Category models.Category
}
