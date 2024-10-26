package HTTPShopProductHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/core/entity/HTTP/shop"
	// "github.com/oogway93/golangArchitecture/internal/core/errors/data/response"
	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
)

type TemplateData struct {
	Categories []models.Category
	Product    productsHTTP.Product
}

func (h *HTTPProductHandler) GetAll(c *gin.Context) {
	// categoryID := c.Param("category")
	// _, result := h.service.GetAll("HTTP", categoryID)

	// webResponse := response.WebResponse{
	// 	Code:   http.StatusOK,
	// 	Status: "Ok",
	// 	Data:   result,
	// }

	// c.Header("Content-Type", "application/json")
	// c.JSON(http.StatusOK, webResponse)
}

func (h *HTTPProductHandler) Get(c *gin.Context) {
	categoryID := c.Param("category")
	productID := c.Param("product")
	resultCategoryModel, _ := h.serviceCategory.Get("HTTP", categoryID)
	resultProductsModel, _ := h.serviceProduct.Get("HTTP", categoryID, productID)

	product := productsHTTP.Product{
		Category: resultCategoryModel,
		Product:  resultProductsModel,
	}
	templateData := &TemplateData{Product: product}
	c.HTML(http.StatusOK, "product", templateData)
}
