package handlerShopOrder

import (
	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/service"
)


type Handler struct {
	service *service.Service
}

func NewOrderShopHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ShopOrderHandlerRoutes(apiRoutes *gin.RouterGroup) *gin.RouterGroup {
	order := apiRoutes.Group("/order")
	{
		order.GET("/", h.GetAll)
		order.POST("/", h.Create)
		// order.GET("/:category", nil)
		// order.PUT("/:category", nil)
		// order.DELETE("/:category", nil)
	}
	
	return order
}

