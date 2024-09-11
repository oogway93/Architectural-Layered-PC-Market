package serverHTTP

import (
	"github.com/gin-gonic/gin"
	handlerAuth "github.com/oogway93/golangArchitecture/internal/core/server/serverHTTP/handler/auth"
	handlerShopCategory "github.com/oogway93/golangArchitecture/internal/core/server/serverHTTP/handler/shop/category"
	handlerShopOrder "github.com/oogway93/golangArchitecture/internal/core/server/serverHTTP/handler/shop/order"
	handlerShopProduct "github.com/oogway93/golangArchitecture/internal/core/server/serverHTTP/handler/shop/product"
	handlerUser "github.com/oogway93/golangArchitecture/internal/core/server/serverHTTP/handler/user"
	"github.com/oogway93/golangArchitecture/internal/core/service"
)

func SetupRouter(
	ServiceCategory service.ServiceCategory,
	ServiceProduct service.ServiceProduct,
	ServiceOrder service.ServiceOrder,
	ServiceUser service.ServiceUser,
	ServiceAuth service.ServiceAuth,
) *gin.Engine {
	router := gin.Default()
	apiRoutes := router.Group("/api", UserIdentity)

	registerShopCategoryRoutes(ServiceCategory, apiRoutes)
	registerShopProductRoutes(ServiceProduct, apiRoutes)
	registerOrderRoutes(ServiceOrder, apiRoutes)
	registerUserRoutes(ServiceUser, router)
	registerAuthRoutes(ServiceAuth, router)

	return router
}

func registerShopCategoryRoutes(service service.ServiceCategory, apiRoutes *gin.RouterGroup) {
	handlerShopCategory.NewCategoryShopHandler(service).ShopCategoryHandlerRoutes(apiRoutes)
}

func registerShopProductRoutes(service service.ServiceProduct, apiRoutes *gin.RouterGroup) {
	handlerShopProduct.NewProductShopHandler(service).ShopProductHandlerRoutes(apiRoutes)
}

func registerOrderRoutes(service service.ServiceOrder, router *gin.RouterGroup) {
	handlerShopOrder.NewOrderShopHandler(service).ShopOrderHandlerRoutes(router)
}

func registerUserRoutes(service service.ServiceUser, router *gin.Engine) {
	handlerUser.NewUserHandler(service).UserHandlerRoutes(router)
}

func registerAuthRoutes(service service.ServiceAuth, router *gin.Engine) {
	handlerAuth.NewAuthHandler(service).AuthHandlerRoutes(router)
}
