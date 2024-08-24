package handlerUser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/entity/user"

	// "github.com/oogway93/golangArchitecture/internal/errors"
	"github.com/oogway93/golangArchitecture/internal/errors/data/response"
	// "github.com/oogway93/golangArchitecture/internal/service"
)

func (h *Handler) Create(c *gin.Context) {
	var newUser user.User

	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	h.service.ServiceUser.Create(&newUser)

	webResponse := response.WebResponse{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   nil,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
