package handlerShopProduct

import (
	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/service"
)

type ProductHandler struct {
	service service.ServiceProduct
}

func NewProductShopHandler(service service.ServiceProduct) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) ShopProductHandlerRoutes(apiRoutes *gin.RouterGroup) *gin.RouterGroup {
	product := apiRoutes.Group("/categories/:category/products")
	{
		product.GET("/", h.GetAll)
		product.POST("/", h.Create)

		product.GET("/:product", h.Get)
		product.PUT("/:product", h.Update)
		product.DELETE("/:product", h.Delete)
	}
	
	return product
}
