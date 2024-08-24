package http

import (
	"github.com/gin-gonic/gin"
	handlerShop "github.com/oogway93/golangArchitecture/internal/server/http/handler/shop"
	handlerUser "github.com/oogway93/golangArchitecture/internal/server/http/handler/user"
	"github.com/oogway93/golangArchitecture/internal/service"
)

func SetupRouter(service *service.Service) *gin.Engine {
	router := gin.Default()
	apiRoutes := router.Group("/api")

	registerShopRoutes(service, apiRoutes)
	registerUserRoutes(service, apiRoutes)

	return router
}

func registerShopRoutes(service *service.Service, apiRoutes *gin.RouterGroup) {
	handlerShop.NewCategoryShopHandler(service).ShopHandlerRoutes(apiRoutes)
}

func registerUserRoutes(service *service.Service, apiRoutes *gin.RouterGroup) {
	handlerUser.NewUserHandler(service).UserHandlerRoutes(apiRoutes)

}
