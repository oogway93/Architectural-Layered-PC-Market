package APIShopProducthandler

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/core/entity/API/shop"
	"github.com/oogway93/golangArchitecture/internal/core/errors/data/response"
)

func (h *ProductHandler) Create(c *gin.Context) {
	var newProduct productsAPI.Product

	categoryID := c.Param("category")

	err := c.BindJSON(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	h.service.Create(categoryID, &newProduct)

	webResponse := response.WebResponse{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   nil,
	}

	c.Header("Content-Type", "application/json")
	c.SecureJSON(http.StatusCreated, webResponse)
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	categoryID := c.Param("category")
	_, result := h.service.GetAll("API", categoryID)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	}

	c.Header("Content-Type", "application/json")
	c.SecureJSON(http.StatusOK, webResponse)
}

func (h *ProductHandler) Get(c *gin.Context) {
	categoryID := c.Param("category")
	productID := c.Param("product")
	if strings.Contains(categoryID, "'") || strings.Contains(categoryID, "-") ||
		strings.Contains(categoryID, "|") || strings.Contains(productID, "'") ||
		strings.Contains(productID, "-") || strings.Contains(productID, "|") {
		slog.Warn("Might Be SQL Injection Attack", "from", c.Request.Host)
		c.SecureJSON(http.StatusBadRequest, response.WebResponse{http.StatusBadRequest, "Wrong Category ID", nil})
		return
	}
	_, result := h.service.Get("API", categoryID, productID)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	}

	c.Header("Content-Type", "application/json")
	c.SecureJSON(http.StatusOK, webResponse)
}
func (h *ProductHandler) Delete(c *gin.Context) {
	categoryID := c.Param("category")
	productID := c.Param("product")
	result := h.service.Delete(categoryID, productID)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "DELETE method doesn't work",
		})
		return
	}
	c.SecureJSON(http.StatusOK, gin.H{
		"message": "Category DELETED successfully",
	})
}
func (h *ProductHandler) Update(c *gin.Context) {
	var newProduct productsAPI.Product
	categoryID := c.Param("category")
	productID := c.Param("product")
	err := c.BindJSON(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	err = h.service.Update(categoryID, productID, &newProduct)
	if err != nil {
		slog.Warn("Errors in Update handler", "error", err.Error())
	}
	c.SecureJSON(http.StatusCreated, gin.H{
		"message": "Category UPDATED successfully",
	})
}
