package handlerShopOrder

import (
	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/service"
)

type OrderHandler struct {
	service service.ServiceOrder
}

func NewOrderShopHandler(service service.ServiceOrder) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) ShopOrderHandlerRoutes(apiRoutes *gin.RouterGroup) *gin.RouterGroup {
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
