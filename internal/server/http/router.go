package HTTP

import (
	"github.com/gin-gonic/gin"
	handlerShopCategory "github.com/oogway93/golangArchitecture/internal/server/http/handler/shop/category"
	handlerShopProduct "github.com/oogway93/golangArchitecture/internal/server/http/handler/shop/product"
	handlerUser "github.com/oogway93/golangArchitecture/internal/server/http/handler/user"
	"github.com/oogway93/golangArchitecture/internal/service"
)

func SetupRouter(service *service.Service) *gin.Engine {
	router := gin.Default()
	apiRoutes := router.Group("/api")

	registerShopCategoryRoutes(service, apiRoutes)
	registerShopProductRoutes(service, apiRoutes)
	registerUserRoutes(service, apiRoutes)

	return router
}

func registerShopCategoryRoutes(service *service.Service, apiRoutes *gin.RouterGroup) {
	handlerShopCategory.NewCategoryShopHandler(service).ShopCategoryHandlerRoutes(apiRoutes)
}

func registerShopProductRoutes(service *service.Service, apiRoutes *gin.RouterGroup) {
	handlerShopProduct.NewProductShopHandler(service).ShopProductHandlerRoutes(apiRoutes)
}

func registerUserRoutes(service *service.Service, apiRoutes *gin.RouterGroup) {
	handlerUser.NewUserHandler(service).UserHandlerRoutes(apiRoutes)
}
