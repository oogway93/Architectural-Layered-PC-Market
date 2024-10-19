package HTTPShopCategoryHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/core/entity/products"
	"github.com/oogway93/golangArchitecture/internal/core/errors/data/response"
)

type TemplateData struct {
	Categories []products.Category
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
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
