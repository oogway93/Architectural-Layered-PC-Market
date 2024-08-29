package handlerShopProduct

import (
	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewProductShopHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ShopProductHandlerRoutes(apiRoutes *gin.RouterGroup) *gin.RouterGroup {
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
