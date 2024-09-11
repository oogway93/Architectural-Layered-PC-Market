package serverHTTP

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/oogway93/golangArchitecture/internal/core/service"
)

type Handler struct {
	service *service.Service
}

func NewMiddlewareHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// FIXME: сделать как-нибудь, чтоб можно было проверить через бд есть ли такой логин,
// иначе опрокинуть ошибку. Проблема просто в том, что как создать структуру Handler, из которой буду вызывать сервис->репозиторий.
func UserIdentity(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := authToken[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	// userId, ok := claims["id"].(uint)
	// if !ok {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
	// 	c.AbortWithStatus(http.StatusUnauthorized)
	// 	return
	// }

	userLogin, ok := claims["login"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login in token"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if float64(time.Now().Unix()) > claims["expiration"].(float64) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// var user models.User
	// initializers.DB.Where("ID=?", claims["id"]).Find(&user)
	// user := h.service.ServiceUser.Get(claims["login"].(string))
	// if user["id"] == 0 {
	// 	c.AbortWithStatus(http.StatusUnauthorized)
	// 	return
	// }

	// c.Set("currentUserID", userId)
	c.Set("currentUserLogin", userLogin)

	c.Next()
}
