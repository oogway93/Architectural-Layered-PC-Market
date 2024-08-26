package handlerShop

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/errors/data/response"
)

func (h *Handler) Create(c *gin.Context) {
	var newCategory products.Category

	err := c.BindJSON(&newCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	h.service.ServiceCategory.Create(&newCategory)

	webResponse := response.WebResponse{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   nil,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, webResponse)
}

func (h *Handler) GetAll(c *gin.Context) {
	result := h.service.ServiceCategory.GetAll()

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

func (h *Handler) Get(c *gin.Context) {
	categoryID := c.Param("category")
	result := h.service.ServiceCategory.Get(categoryID)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   gin.H{"category_name": result},
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
func (h *Handler) Delete(c *gin.Context) {}
func (h *Handler) Update(c *gin.Context) {}
 