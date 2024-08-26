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

func (h *Handler) ShopCategoryHandlerRoutes(apiRoutes *gin.RouterGroup) *gin.RouterGroup {
	category := apiRoutes.Group("/category")
	{
		category.GET("/", nil)
		category.POST("/", h.Create)
		category.GET("/:categoryID", nil)
		category.PUT("/:categoryID", nil)
		category.DELETE("/:categoryID", nil)
	}
	
	return category
}
