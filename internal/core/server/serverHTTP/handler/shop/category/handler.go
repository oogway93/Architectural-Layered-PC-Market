package HTTPShopCategoryHandler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/oogway93/golangArchitecture/internal/core/entity/products"
	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
	"github.com/shopspring/decimal"
	// "github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
)

type TemplateData struct {
	Categories       []products.Category
	CategoryProducts products.CategoryProducts
}

func (h *HTTPCategoryHandler) GetAll(c *gin.Context) {
	var categories []products.Category
	result := h.service.GetAll()
	for _, category := range result {
		category_name := products.Category{
			CategoryName: category["category_name"].(string),
		}

		categories = append(categories, category_name)
	}
	templateData := &TemplateData{Categories: categories}
	c.HTML(http.StatusOK, "categories.html", templateData)
}

func (h *HTTPCategoryHandler) Get(c *gin.Context) {
	categoryID := c.Param("category")
	result := h.service.Get(categoryID)
	var categoryProduct []models.Product
	if sliceOfInterfaces, ok := result["products"].([]interface{}); ok {
		for _, productInterface := range sliceOfInterfaces {
			if productMap, ok := productInterface.(map[string]interface{}); ok {
				price, err := decimal.NewFromString(productMap["price"].(string))
				if err != nil {
					slog.Info("Incorrect type for Price", "error", err)
				}
				product := models.Product{
					UUID:        uuid.MustParse(productMap["uuid"].(string)),
					ProductName: productMap["product_name"].(string),
					Description: productMap["description"].(string),
					Price:       price,
					Category:    models.Category{CategoryName: result["category_name"].(string)},
				}
				categoryProduct = append(categoryProduct, product)
			} else {
				slog.Warn("Error: Unexpected type in products slice. Expected map[string]interface{}, got %T", productInterface)
			}
		}
	} else {
		slog.Info("Error: Unexpected type for products. Expected []interface{}, got %T", result["products"])
	}
	categoryProducts := products.CategoryProducts{
		CategoryName: result["category_name"].(string),
		Products:     categoryProduct,
	}
	templateData := &TemplateData{CategoryProducts: categoryProducts}
	c.HTML(http.StatusOK, "category.html", templateData)
}
