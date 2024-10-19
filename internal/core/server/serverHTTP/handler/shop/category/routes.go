package HTTPShopCategoryHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/core/service"
)

type HTTPCategoryHandler struct {
	service service.ServiceCategory
}

func NewHTTPCategoryShopHandler(service service.ServiceCategory) *HTTPCategoryHandler {
	return &HTTPCategoryHandler{
		service: service,
	}
}

func (h *HTTPCategoryHandler) HTTPShopCategoryHandlerRoutes(httpRoutes *gin.RouterGroup) *gin.RouterGroup {
	category := httpRoutes.Group("/category")
	{
		category.GET("/", h.GetAll)
		// category.POST("/", h.Create)
		category.GET("/:category", h.Get)
		// category.PUT("/:category", h.Update)
		// category.DELETE("/:category", h.Delete)
	}

	return category
}
