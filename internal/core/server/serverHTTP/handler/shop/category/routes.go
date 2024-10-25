package HTTPShopCategoryHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/core/service"
)

type HTTPCategoryHandler struct {
	serviceCategory service.ServiceCategory
	serviceProduct service.ServiceProduct
}

func NewHTTPCategoryShopHandler(serviceCategory service.ServiceCategory, serviceProduct service.ServiceProduct) *HTTPCategoryHandler {
	return &HTTPCategoryHandler{
		serviceCategory: serviceCategory,
		serviceProduct: serviceProduct,
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
