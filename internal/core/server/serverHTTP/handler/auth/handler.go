package handlerAuth

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/oogway93/golangArchitecture/internal/core/entity/user"
	"github.com/oogway93/golangArchitecture/internal/core/errors/data/response"
)

func (h *AuthHandler) Login(c *gin.Context) {
	var authInput user.AuthInput

	if err := c.BindJSON(&authInput); err != nil {
		log.Fatalf("Error LOGIN handler: %v", err.Error())
	}

	result := h.service.Login(&authInput)

	c.Header("Content-Type", "application/json")
	if !result {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid password",
		})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login":      authInput.Login,
		"expiration": time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data: gin.H{
			"token": token,
		},
	}

	c.JSON(http.StatusOK, webResponse)
}
