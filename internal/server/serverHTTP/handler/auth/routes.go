package handlerAuth

import (
	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/service"
)

type AuthHandler struct {
	service service.ServiceAuth
}

func NewAuthHandler(service service.ServiceAuth) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) AuthHandlerRoutes(router *gin.Engine) *gin.RouterGroup {
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
