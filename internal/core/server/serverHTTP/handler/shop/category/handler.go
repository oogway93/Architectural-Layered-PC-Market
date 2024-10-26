package HTTPShopCategoryHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/core/entity/HTTP/shop"
	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
)

type TemplateData struct {
	Categories       []models.Category
	CategoryProducts productsHTTP.CategoryProducts
}

func (h *HTTPCategoryHandler) GetAll(c *gin.Context) {
	resultCategoryModel, _ := h.serviceCategory.GetAll("HTTP")
	templateData := &TemplateData{Categories: resultCategoryModel}
	c.HTML(http.StatusOK, "categories", templateData)
}

func (h *HTTPCategoryHandler) Get(c *gin.Context) {
	categoryID := c.Param("category")
	resultCategory, _ := h.serviceCategory.Get("HTTP", categoryID)
	resultProductsModels, _ := h.serviceProduct.GetAll("HTTP", categoryID)

	categoryProducts := productsHTTP.CategoryProducts{
		Category: resultCategory,
		Products: resultProductsModels,
	}
	templateData := &TemplateData{CategoryProducts: categoryProducts}
	c.HTML(http.StatusOK, "categoryWithProducts", templateData)
}
