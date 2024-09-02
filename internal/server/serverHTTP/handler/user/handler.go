package handlerUser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/entity/user"

	"github.com/oogway93/golangArchitecture/internal/errors/data/response"
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
	c.JSON(http.StatusCreated, webResponse)
}

func (h *Handler) GetAll(c *gin.Context) {
	result := h.service.ServiceUser.GetAll()

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

func (h *Handler) Update(c *gin.Context) {
	var newUser user.UserUpdated

	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	loginID := c.Param("login")

	h.service.ServiceUser.Update(loginID, &newUser)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User UPDATED successfully",
	})
}

func (h *Handler) Delete(c *gin.Context) {
	loginID := c.Param("login")
	result := h.service.ServiceUser.Delete(loginID)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "DELETE method doesn't work",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User DELETED successfully",
	})
}
