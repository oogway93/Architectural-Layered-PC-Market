package productsHTTP

import "github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"

type CategoryProducts struct {
	Category models.Category
	Products []models.Product
}

type Categories struct {
	Categories []models.Category
}
