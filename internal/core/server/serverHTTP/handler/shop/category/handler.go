package HTTPShopCategoryHandler

import (
	"log"
	// "log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
	"github.com/oogway93/golangArchitecture/internal/core/entity/products"
	// "github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
	// "github.com/shopspring/decimal"
	// "github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
)

type TemplateData struct {
	Categories       []products.Category
	CategoryProducts products.CategoryProducts
}

func (h *HTTPCategoryHandler) GetAll(c *gin.Context) {
	var categories []products.Category
	result := h.serviceCategory.GetAll()
	log.Println(result)
	for _, category := range result {
		category_name := products.Category{
			CategoryName: category["category_name"].(string),
		}

		categories = append(categories, category_name)
	}
	templateData := &TemplateData{Categories: categories}
	c.HTML(http.StatusOK, "categories", templateData)

}

func (h *HTTPCategoryHandler) Get(c *gin.Context) {
	categoryID := c.Param("category")
	resultCategory := h.serviceCategory.Get(categoryID)
	resultProductsModels, _ := h.serviceProduct.GetAll("HTTP", categoryID)
	// var categoryProduct []models.Product
	// for _, product := range resultProducts {
	// 	// price, err := decimal.NewFromString(resultCategory["products"].(map[string]interface{}))
	// 	// if err != nil {
	// 	// 	slog.Info("Incorrect type for Price", "error", err)
	// 	// }
	// 	product := models.Product{
	// 		UUID:        uuid.MustParse(product["uuid"].(string)),
	// 		ProductName: product["product_name"].(string),
	// 		Description: product["description"].(string),
	// 		Price:       product["price"].(decimal.Decimal),
	// 		Category:    models.Category{CategoryName: resultCategory["category_name"].(string)},
	// 	}
	// 	categoryProduct = append(categoryProduct, product)
	// }

	categoryProducts := products.CategoryProducts{
		CategoryName: resultCategory["category_name"].(string),
		Products:     resultProductsModels,
	}
	templateData := &TemplateData{CategoryProducts: categoryProducts}
	c.HTML(http.StatusOK, "category", templateData)

}
