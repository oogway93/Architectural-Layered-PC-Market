package handlerShopOrder

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/errors/data/response"
	"github.com/oogway93/golangArchitecture/internal/utils"
)

func (h *Handler) Create(c *gin.Context) {
	var order products.Order
	
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	userID, err := utils.GetUserID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "get userID is incorrect",
		})
	}

	h.service.ServiceOrder.Create(userID, &order)
	
	webResponse := response.WebResponse{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   nil,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, webResponse)
}

func (h *Handler) GetAll(c *gin.Context) {
	userID, err := utils.GetUserID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "get userID is incorrect",
		})
	}
	result := h.service.ServiceOrder.GetAll(userID)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
