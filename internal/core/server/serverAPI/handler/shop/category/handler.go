package handlerShopCategory

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/core/entity/API/shop"
	"github.com/oogway93/golangArchitecture/internal/core/errors/data/response"
)

func (h *CategoryHandler) Create(c *gin.Context) {
	var newCategory productsAPI.Category

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
	c.SecureJSON(http.StatusCreated, webResponse)
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	_, result := h.service.GetAll("API")
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	}

	c.Header("Content-Type", "application/json")
	c.SecureJSON(http.StatusOK, webResponse)
}

func (h *CategoryHandler) Get(c *gin.Context) {
	categoryID := c.Param("category")
	c.Header("Content-Type", "application/json")
	if strings.Contains(categoryID, "'") || strings.Contains(categoryID, "-") || strings.Contains(categoryID, "|") {
		slog.Warn("Might Be SQL Injection Attack", "from", c.Request.Host)
		c.SecureJSON(http.StatusBadRequest, response.WebResponse{http.StatusBadRequest, "Wrong Category ID", nil})
		return
	}
	_, result := h.service.Get("API", categoryID)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	}

	c.SecureJSON(http.StatusOK, webResponse)
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
	c.SecureJSON(http.StatusOK, gin.H{
		"message": "Category DELETED successfully",
	})
}
func (h *CategoryHandler) Update(c *gin.Context) {
	var newCategory productsAPI.Category
	categoryID := c.Param("category")
	err := c.BindJSON(&newCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	h.service.Update(categoryID, &newCategory)
	c.SecureJSON(http.StatusCreated, gin.H{
		"message": "Category UPDATED successfully",
	})
}
