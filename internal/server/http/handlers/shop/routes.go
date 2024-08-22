package handlersShop

import "github.com/gin-gonic/gin"

func HandlerRoutes(c *gin.Context) {
	r:= gin.Default()
	
	api := r.Group("/api", nil)
	{
		productList := r.Group("/productList")
		{
			productList.GET("/", nil)
			productList.POST("/", nil)
			productList.GET("/:id", nil)
			productList.PUT("/:id", nil)
			productList.DELETE("/:id", nil)
		}
	}
}
