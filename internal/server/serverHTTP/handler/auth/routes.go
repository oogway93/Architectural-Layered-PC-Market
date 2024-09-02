package handlerAuth

import (
	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewAuthHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AuthHandlerRoutes(router *gin.Engine) *gin.RouterGroup {
	auth := router.Group("/user/auth")
	{
		auth.POST("/login", h.Login)
		// auth.POST("/", h.Create)
		// auth.GET("/:login", nil)
		// auth.PUT("/:login", h.Update)
		// auth.DELETE("/:login", h.Delete)
	}

	return auth
}
