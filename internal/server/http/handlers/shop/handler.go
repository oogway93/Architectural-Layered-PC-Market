package handlerShop

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloWorld struct {
	Msg string `json:"msg"`
}
func HelloWorldCon(c *gin.Context) {
	d := &HelloWorld{"Hello world"}
	c.JSON(http.StatusAccepted, d)
}	


