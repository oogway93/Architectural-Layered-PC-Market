package server

import (
	"github.com/gin-gonic/gin"
	APIHandlerShopCategory "github.com/oogway93/golangArchitecture/internal/core/server/serverAPI/handler/shop/category"
	HTTPHandlerShopCategory "github.com/oogway93/golangArchitecture/internal/core/server/serverHTTP/handler/shop/category"
	APIHandlerShopOrder "github.com/oogway93/golangArchitecture/internal/core/server/serverAPI/handler/shop/order"
	APIHandlerShopProduct "github.com/oogway93/golangArchitecture/internal/core/server/serverAPI/handler/shop/product"
	APIHandlerUser "github.com/oogway93/golangArchitecture/internal/core/server/serverAPI/handler/user"
	APIHandlerAuth "github.com/oogway93/golangArchitecture/internal/core/server/serverAPI/handler/auth"
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
	router.LoadHTMLGlob("internal/core/server/serverHTTP/static/templates/*")
	httpRoutes := router.Group("/")
	registerAPIRoutes(apiRoutes, httpRoutes, ServiceCategory, ServiceProduct, ServiceOrder, ServiceUser, ServiceAuth)
	registerHTTPRoutes(httpRoutes, ServiceCategory)
	return router
}

func registerHTTPRoutes(
	httpRoutes *gin.RouterGroup,
	ServiceCategory service.ServiceCategory,
	// ServiceProduct service.ServiceProduct,
	// ServiceOrder service.ServiceOrder,
	// ServiceUser service.ServiceUser,
	// ServiceAuth service.ServiceAuth,
) {
	registerHTTPShopCategoryRoutes(ServiceCategory, httpRoutes)
}

func registerAPIRoutes(
	apiRoutes *gin.RouterGroup,
	httpRoutes *gin.RouterGroup,
	ServiceCategory service.ServiceCategory,
	ServiceProduct service.ServiceProduct,
	ServiceOrder service.ServiceOrder,
	ServiceUser service.ServiceUser,
	ServiceAuth service.ServiceAuth,
) {
	registerAPIShopCategoryRoutes(ServiceCategory, apiRoutes)
	registerAPIShopProductRoutes(ServiceProduct, apiRoutes)
	registerAPIOrderRoutes(ServiceOrder, apiRoutes)
	registerAPIUserRoutes(ServiceUser, httpRoutes)
	registerAPIAuthRoutes(ServiceAuth, httpRoutes)
}

func registerHTTPShopCategoryRoutes(service service.ServiceCategory, router *gin.RouterGroup) {
	HTTPHandlerShopCategory.NewHTTPCategoryShopHandler(service).HTTPShopCategoryHandlerRoutes(router)
}


func registerAPIShopCategoryRoutes(service service.ServiceCategory, router *gin.RouterGroup) {
	APIHandlerShopCategory.NewCategoryShopHandler(service).ShopCategoryHandlerRoutes(router)
}

func registerAPIShopProductRoutes(service service.ServiceProduct, router *gin.RouterGroup) {
	APIHandlerShopProduct.NewProductShopHandler(service).ShopProductHandlerRoutes(router)
}

func registerAPIOrderRoutes(service service.ServiceOrder, router *gin.RouterGroup) {
	APIHandlerShopOrder.NewOrderShopHandler(service).ShopOrderHandlerRoutes(router)
}

func registerAPIUserRoutes(service service.ServiceUser, router *gin.RouterGroup) {
	APIHandlerUser.NewUserHandler(service).UserHandlerRoutes(router)
}

func registerAPIAuthRoutes(service service.ServiceAuth, router *gin.RouterGroup) {
	APIHandlerAuth.NewAuthHandler(service).AuthHandlerRoutes(router)
}
