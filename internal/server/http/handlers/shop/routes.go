package handlerShop

import "github.com/gin-gonic/gin"

// type Handler struct {
// 	services *service.Service
// }

// func NewHandler(services *service.Service) *Handler {
// 	return &Handler{services: services}
// }

func HandlerRoutes() *gin.Engine{
	r:= gin.Default()
	
	api := r.Group("/api")
	{
		productList := api.Group("/productList")
		{
			productList.GET("/", HelloWorldCon)
			productList.POST("/", nil)
			productList.GET("/:id", nil)
			productList.PUT("/:id", nil)
			productList.DELETE("/:id", nil)
		}
	}
	return r
}
