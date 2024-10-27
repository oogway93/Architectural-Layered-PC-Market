package handlerAuth

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/oogway93/golangArchitecture/internal/core/entity/API/user"
	"github.com/oogway93/golangArchitecture/internal/core/errors/data/response"
)

func (h *AuthHandler) Login(c *gin.Context) {
	var authInput user.AuthInput

	if err := c.BindJSON(&authInput); err != nil {
		slog.Warn("Error LOGIN handler", "error", err.Error())
	}

	result := h.service.Login(&authInput)

	c.Header("Content-Type", "application/json")
	if !result {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid password",
		})
		return
	}

	token := GenerateToken(c, authInput.Login)
	if token == "" {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
		return
	}
	SetJWTToken(c, token)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data: gin.H{
			"token": token,
		},
	}
	c.SecureJSON(http.StatusOK, webResponse)
}

func GenerateToken(c *gin.Context, login string) string {
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login":      login,
		"expiration": time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.HTML(http.StatusBadRequest, "login", gin.H{"ErrorMessage": "Failed to generate token"})
		return ""
	}
	return token
}

func SetJWTToken(c *gin.Context, token string) {
	expire := time.Now().Add(time.Hour * 24)
	c.SetCookie("jwt", token, int(expire.Unix()),
		"/", "", true, true,
	)
}

func GetJWTToken(c *gin.Context) (string, error) {
	cookie, err := c.Request.Cookie("jwt")
	if err != nil || cookie == nil {
		return "", errors.New("no JWT token found")
	}
	return cookie.Value, nil
}
