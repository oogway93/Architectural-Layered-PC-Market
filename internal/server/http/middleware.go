package serverHTTP

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("userId")
	if !ok {
		return 0, errors.New("Cannot get user's ID")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("id is of invalid type")
	}

	return idInt, nil
}	

