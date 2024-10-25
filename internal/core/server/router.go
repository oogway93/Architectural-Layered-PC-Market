package server

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	APIHandlerAuth "github.com/oogway93/golangArchitecture/internal/core/server/serverAPI/handler/auth"
	APIHandlerShopCategory "github.com/oogway93/golangArchitecture/internal/core/server/serverAPI/handler/shop/category"
	APIHandlerShopOrder "github.com/oogway93/golangArchitecture/internal/core/server/serverAPI/handler/shop/order"
	APIHandlerShopProduct "github.com/oogway93/golangArchitecture/internal/core/server/serverAPI/handler/shop/product"
	APIHandlerUser "github.com/oogway93/golangArchitecture/internal/core/server/serverAPI/handler/user"
	HTTPHandlerShopCategory "github.com/oogway93/golangArchitecture/internal/core/server/serverHTTP/handler/shop/category"
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
	router.HTMLRender = ginview.New(goview.Config{
		Root:      "internal/core/server/serverHTTP/static/templates/shop",
		Extension: ".tmpl",
		Master:    "base",
		Partials:  nil,
		Funcs: nil,
		DisableCache: true,
	})
	apiRoutes := router.Group("/api", UserIdentity)
	httpRoutes := router.Group("/")
	registerHTTPRoutes(httpRoutes, ServiceCategory, ServiceProduct)
	registerAPIRoutes(apiRoutes, httpRoutes, ServiceCategory, ServiceProduct, ServiceOrder, ServiceUser, ServiceAuth)
	return router
}

func registerHTTPRoutes(
	httpRoutes *gin.RouterGroup,
	ServiceCategory service.ServiceCategory,
	ServiceProduct service.ServiceProduct,
	// ServiceOrder service.ServiceOrder,
	// ServiceUser service.ServiceUser,
	// ServiceAuth service.ServiceAuth,
) {
	registerHTTPShopCategoryRoutes(ServiceCategory, ServiceProduct, httpRoutes)
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

func registerHTTPShopCategoryRoutes(serviceCategory service.ServiceCategory, serviceProduct service.ServiceProduct, router *gin.RouterGroup) {
	HTTPHandlerShopCategory.NewHTTPCategoryShopHandler(serviceCategory, serviceProduct).HTTPShopCategoryHandlerRoutes(router)
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
