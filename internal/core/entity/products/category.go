package products

import "github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"

type Category struct {
	CategoryName string `json:"category_name"`
}

type CategoryProducts struct {
	CategoryName string `json:"category_name"`
	Products    []models.Product
}
