package handlerShopCategory

import (
	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/service"
)

type CategoryHandler struct {
	service service.ServiceCategory
}

func NewCategoryShopHandler(service service.ServiceCategory) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (h *CategoryHandler) ShopCategoryHandlerRoutes(apiRoutes *gin.RouterGroup) *gin.RouterGroup {
	category := apiRoutes.Group("/categories")
	{
		category.GET("/", h.GetAll)
		category.POST("/", h.Create)
		category.GET("/:category", h.Get)
		category.PUT("/:category", h.Update)
		category.DELETE("/:category", h.Delete)
	}
	
	return category
}
