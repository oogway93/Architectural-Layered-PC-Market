package handlerShop

import (
	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewCategoryShopHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ShopHandlerRoutes(apiRoutes *gin.RouterGroup) *gin.RouterGroup {
	productList := apiRoutes.Group("/productList")
	{
		productList.GET("/", h.HelloWorldCon)
		productList.POST("/", nil)
		productList.GET("/:id", nil)
		productList.PUT("/:id", nil)
		productList.DELETE("/:id", nil)
	}
	
	return productList
}
