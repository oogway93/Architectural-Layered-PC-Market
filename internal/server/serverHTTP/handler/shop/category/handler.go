package handlerShopCategory

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/errors/data/response"
)

func (h *CategoryHandler) Create(c *gin.Context) {
	var newCategory products.Category

	err := c.BindJSON(&newCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	h.service.Create(&newCategory)

	webResponse := response.WebResponse{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   nil,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, webResponse)
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	result := h.service.GetAll()

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

func (h *CategoryHandler) Get(c *gin.Context) {
	categoryID := c.Param("category")
	result := h.service.Get(categoryID)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
func (h *CategoryHandler) Delete(c *gin.Context) {
	categoryID := c.Param("category")
	result := h.service.Delete(categoryID)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "DELETE method doesn't work",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Category DELETED successfully",
	})
}
func (h *CategoryHandler) Update(c *gin.Context) {
	var newCategory products.Category
	categoryID := c.Param("category")
	err := c.BindJSON(&newCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	h.service.Update(categoryID, &newCategory)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Category UPDATED successfully",
	})
}
