package HTTP

import (
	"github.com/gin-gonic/gin"
	handlerAuth "github.com/oogway93/golangArchitecture/internal/server/http/handler/auth"
	handlerShopCategory "github.com/oogway93/golangArchitecture/internal/server/http/handler/shop/category"
	handlerShopProduct "github.com/oogway93/golangArchitecture/internal/server/http/handler/shop/product"
	handlerUser "github.com/oogway93/golangArchitecture/internal/server/http/handler/user"
	"github.com/oogway93/golangArchitecture/internal/service"
)


func SetupRouter(service *service.Service) *gin.Engine {
	router := gin.Default()
	apiRoutes := router.Group("/api", UserIdentity)

	registerShopCategoryRoutes(service, apiRoutes)
	registerShopProductRoutes(service, apiRoutes)
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

func registerUserRoutes(service *service.Service, router *gin.Engine) {
	handlerUser.NewUserHandler(service).UserHandlerRoutes(router)
}

func registerAuthRoutes(service *service.Service, router *gin.Engine) {
	handlerAuth.NewAuthHandler(service).AuthHandlerRoutes(router)
}

// func registerMiddlewareRoutes(service *service.Service) {
// 	NewMiddlewareHandler(service)
// }
