package handlerUser

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type HelloWorld struct {
	Msg string `json:"msg"`
}
func (h *Handler) HelloWorldCon(c *gin.Context) {
	d := &HelloWorld{"Hello world from /User"}
	c.JSON(http.StatusAccepted, d)
}	


