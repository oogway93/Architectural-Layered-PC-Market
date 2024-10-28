package server

import (
	"net/http"
	"text/template"
	"time"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/adapter/config"
	APIHandlerAuth "github.com/oogway93/golangArchitecture/internal/core/server/serverAPI/handler/auth"
	APIHandlerShopCategory "github.com/oogway93/golangArchitecture/internal/core/server/serverAPI/handler/shop/category"
	APIHandlerShopOrder "github.com/oogway93/golangArchitecture/internal/core/server/serverAPI/handler/shop/order"
	APIHandlerShopProduct "github.com/oogway93/golangArchitecture/internal/core/server/serverAPI/handler/shop/product"
	APIHandlerUser "github.com/oogway93/golangArchitecture/internal/core/server/serverAPI/handler/user"
	HTTPHandlerShopCategory "github.com/oogway93/golangArchitecture/internal/core/server/serverHTTP/handler/shop/category"
	HTTPHandlerShopProduct "github.com/oogway93/golangArchitecture/internal/core/server/serverHTTP/handler/shop/product"
	HTTPAuthHandler "github.com/oogway93/golangArchitecture/internal/core/server/serverHTTP/handler/user"
	"github.com/oogway93/golangArchitecture/internal/core/service"
)

func SetupRouter(
	config *config.HTTP,
	ServiceCategory service.ServiceCategory,
	ServiceProduct service.ServiceProduct,
	ServiceOrder service.ServiceOrder,
	ServiceUser service.ServiceUser,
	ServiceAuth service.ServiceAuth,
) *gin.Engine {
	router := gin.Default()
	router.HTMLRender = ginview.New(goview.Config{
		Root:      config.TemplatesPath,
		Extension: ".html",
		Master:    "base",
		Partials:  []string{"boostrap", "nav"},
		Funcs: template.FuncMap{"humanDate": func(t time.Time) string {
			return t.Format("02 Jan 2006 at 15:04")
		}},
		DisableCache: true,
	})
	router.Use(secureHeaders)
	router.Use(cors.Default())
	apiRoutes := router.Group("/api", UserIdentity)
	httpRoutes := router.Group("/")
	registerHTTPRoutes(httpRoutes, ServiceCategory, ServiceProduct, ServiceAuth, ServiceUser)
	registerAPIRoutes(apiRoutes, ServiceCategory, ServiceProduct, ServiceOrder, ServiceUser, ServiceAuth)
	router.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "home", nil) })
	return router
}

func registerHTTPRoutes(
	httpRoutes *gin.RouterGroup,
	ServiceCategory service.ServiceCategory,
	ServiceProduct service.ServiceProduct,
	ServiceAuth service.ServiceAuth,
	ServiceUser service.ServiceUser,
	// ServiceOrder service.ServiceOrder,
) {
	registerHTTPShopCategoryRoutes(ServiceCategory, ServiceProduct, httpRoutes)
	registerHTTPShopProductRoutes(ServiceCategory, ServiceProduct, httpRoutes)
	registerHTTPAuthRoutes(ServiceAuth, ServiceUser, httpRoutes)
}

func registerAPIRoutes(
	apiRoutes *gin.RouterGroup,
	ServiceCategory service.ServiceCategory,
	ServiceProduct service.ServiceProduct,
	ServiceOrder service.ServiceOrder,
	ServiceUser service.ServiceUser,
	ServiceAuth service.ServiceAuth,
) {
	registerAPIShopCategoryRoutes(ServiceCategory, apiRoutes)
	registerAPIShopProductRoutes(ServiceProduct, apiRoutes)
	registerAPIOrderRoutes(ServiceOrder, apiRoutes)
	registerAPIUserRoutes(ServiceUser, apiRoutes)
	registerAPIAuthRoutes(ServiceAuth, apiRoutes)
}

func registerHTTPShopCategoryRoutes(serviceCategory service.ServiceCategory, serviceProduct service.ServiceProduct, router *gin.RouterGroup) {
	HTTPHandlerShopCategory.NewHTTPCategoryShopHandler(serviceCategory, serviceProduct).HTTPShopCategoryHandlerRoutes(router)
}

func registerHTTPShopProductRoutes(serviceCategory service.ServiceCategory, serviceProduct service.ServiceProduct, router *gin.RouterGroup) {
	HTTPHandlerShopProduct.NewHTTPProductShopHandler(serviceCategory, serviceProduct).HTTPShopProductHandlerRoutes(router)
}

func registerHTTPAuthRoutes(serviceAuth service.ServiceAuth, serviceUser service.ServiceUser, router *gin.RouterGroup) {
	HTTPAuthHandler.NewHTTPAuthHandler(serviceAuth, serviceUser).HTTPAuthHandlerRoutes(router)
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
