package handlerUser

import (
	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/core/service"
)

type UserHandler struct {
	service service.ServiceUser
}

func NewUserHandler(service service.ServiceUser) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) UserHandlerRoutes(router *gin.Engine) *gin.RouterGroup {
	user := router.Group("/user")
	{
		user.GET("/", h.GetAll)
		user.POST("/", h.Create)
		user.GET("/:login", nil)
		user.PUT("/:login", h.Update)
		user.DELETE("/:login", h.Delete)
	}

	return user
}
