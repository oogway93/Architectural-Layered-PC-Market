package HTTPUserHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/core/service"
)

type HTTPAuthHandler struct {
	serviceAuth service.ServiceAuth
	serviceUser service.ServiceUser
}

func NewHTTPAuthHandler(serviceAuth service.ServiceAuth, serviceUser service.ServiceUser) *HTTPAuthHandler {
	return &HTTPAuthHandler{
		serviceAuth: serviceAuth,
		serviceUser: serviceUser,
	}
}

func (h *HTTPAuthHandler) HTTPAuthHandlerRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	auth := router.Group("/user/auth")
	{
		auth.POST("/login", h.Login)
		auth.GET("/login", h.LoginPage)
		auth.POST("/registration", h.Registration)
		auth.GET("/registration", h.RegistrationPage)
	}
	return auth
}
