package HTTPShopProductHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/core/errors/data/response"
)

func (h *HTTPProductHandler) GetAll(c *gin.Context) {
	categoryID := c.Param("category")
	_,result := h.service.GetAll("API", categoryID)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

func (h *HTTPProductHandler) Get(c *gin.Context) {
	categoryID := c.Param("category")
	productID := c.Param("product")
	result := h.service.Get(categoryID, productID)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
