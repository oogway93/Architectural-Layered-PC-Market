package HTTPShopProductHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/core/service"
)

type HTTPProductHandler struct {
	service service.ServiceProduct
}

func NewProductShopHandler(service service.ServiceProduct) *HTTPProductHandler {
	return &HTTPProductHandler{
		service: service,
	}
}

func (h *HTTPProductHandler) HTTPShopProductHandlerRoutes(HTTPRoutes *gin.RouterGroup) *gin.RouterGroup {
	product := HTTPRoutes.Group("/category/:category/products")
	{
		product.GET("/", h.GetAll)
		// product.POST("/", h.Create)

		product.GET("/:product", h.Get)
		// product.PUT("/:product", h.Update)
		// product.DELETE("/:product", h.Delete)
	}

	return product
}
