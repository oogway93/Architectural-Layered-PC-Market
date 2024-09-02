package serverHTTP

import (
	"github.com/gin-gonic/gin"
	handlerAuth "github.com/oogway93/golangArchitecture/internal/server/serverHTTP/handler/auth"
	handlerShopCategory "github.com/oogway93/golangArchitecture/internal/server/serverHTTP/handler/shop/category"
	handlerShopOrder "github.com/oogway93/golangArchitecture/internal/server/serverHTTP/handler/shop/order"
	handlerShopProduct "github.com/oogway93/golangArchitecture/internal/server/serverHTTP/handler/shop/product"
	handlerUser "github.com/oogway93/golangArchitecture/internal/server/serverHTTP/handler/user"
	"github.com/oogway93/golangArchitecture/internal/service"
)


func SetupRouter(service *service.Service) *gin.Engine {
	router := gin.Default()
	apiRoutes := router.Group("/api", UserIdentity)

	registerShopCategoryRoutes(service, apiRoutes)
	registerShopProductRoutes(service, apiRoutes)
	registerOrderRoutes(service, apiRoutes)
	registerUserRoutes(service, router)
	registerAuthRoutes(service, router)

	return router
}

func registerShopCategoryRoutes(service *service.Service, apiRoutes *gin.RouterGroup) {
	handlerShopCategory.NewCategoryShopHandler(service).ShopCategoryHandlerRoutes(apiRoutes)
}

func registerShopProductRoutes(service *service.Service, apiRoutes *gin.RouterGroup) {
	handlerShopProduct.NewProductShopHandler(service).ShopProductHandlerRoutes(apiRoutes)
}

func registerOrderRoutes(service *service.Service, router *gin.RouterGroup) {
	handlerShopOrder.NewOrderShopHandler(service).ShopOrderHandlerRoutes(router)
}

func registerUserRoutes(service *service.Service, router *gin.Engine) {
	handlerUser.NewUserHandler(service).UserHandlerRoutes(router)
}

func registerAuthRoutes(service *service.Service, router *gin.Engine) {
	handlerAuth.NewAuthHandler(service).AuthHandlerRoutes(router)
}
