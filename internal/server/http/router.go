package HTTP

import (
	"github.com/gin-gonic/gin"
	handlerUser "github.com/oogway93/golangArchitecture/internal/server/http/handler/user"
	"github.com/oogway93/golangArchitecture/internal/service"
)

func SetupRouter(service *service.Service) *gin.Engine {
	router := gin.Default()
	apiRoutes := router.Group("/api")

	registerShopCAtegoryRoutes(service, apiRoutes)
	registerUserRoutes(service, apiRoutes)

	return router
}

func registerShopCAtegoryRoutes(service *service.Service, apiRoutes *gin.RouterGroup) {
	handlerShopCategory.NewCategoryShopHandler(service).ShopCategoryHandlerRoutes(apiRoutes)
}

func registerUserRoutes(service *service.Service, apiRoutes *gin.RouterGroup) {
	handlerUser.NewUserHandler(service).UserHandlerRoutes(apiRoutes)

}
