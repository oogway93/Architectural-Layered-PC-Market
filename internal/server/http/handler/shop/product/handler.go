package handlerShopProduct

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/errors/data/response"
)

func (h *Handler) Create(c *gin.Context) {
	var newProduct products.Product

	categoryID := c.Param("category")

	err := c.BindJSON(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	h.service.ServiceProduct.Create(categoryID, &newProduct)

	webResponse := response.WebResponse{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   nil,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, webResponse)
}

func (h *Handler) GetAll(c *gin.Context) {
	categoryID := c.Param("category")
	result := h.service.ServiceProduct.GetAll(categoryID)

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
	productID := c.Param("product")
	result := h.service.ServiceProduct.Get(productID, categoryID)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   gin.H{"category_name": result},
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
func (h *Handler) Delete(c *gin.Context) {
	categoryID := c.Param("category")
	productID := c.Param("product")
	result := h.service.ServiceProduct.Delete(categoryID, productID)
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
func (h *Handler) Update(c *gin.Context) {
	var newProduct products.Product
	categoryID := c.Param("category")
	productID := c.Param("product")
	err := c.BindJSON(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	h.service.ServiceProduct.Update(categoryID, productID, &newProduct)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Category UPDATED successfully",
	})
}
