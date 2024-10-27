package HTTPUserHandler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/core/entity/API/user"
	handlerAuth "github.com/oogway93/golangArchitecture/internal/core/server/serverAPI/handler/auth"
)

func (h *HTTPAuthHandler) Login(c *gin.Context) {
	login := strings.Trim(c.PostForm("login"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	authInputForm := &user.AuthInput{
		Login:    login,
		Password: password,
	}
	result := h.serviceAuth.Login(authInputForm)

	if login == "" || password == "" {
		c.HTML(http.StatusBadRequest, "login", gin.H{"ErrorMessage": "Parameters can't be empty"})
		return
	}

	if !result {
		c.HTML(http.StatusBadRequest, "login", gin.H{"ErrorMessage": "Invalid login's data"})
		return
	} else {
		token := handlerAuth.GenerateToken(c, login)
		if token == "" {
			c.HTML(http.StatusBadRequest, "login", gin.H{"ErrorMessage": "Invalid token"})
			return
		}
		c.SetSameSite(http.SameSiteNoneMode)
		handlerAuth.SetJWTToken(c, token)
		c.Set("currentUserLogin", login)
		c.Redirect(http.StatusFound, "/")
	}
}

func (h *HTTPAuthHandler) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login", nil)
}

func (h *HTTPAuthHandler) RegistrationPage(c *gin.Context) {
	c.HTML(http.StatusOK, "registration", nil)
}

func (h *HTTPAuthHandler) Registration(c *gin.Context) {
	login := strings.Trim(c.PostForm("login"), " ")
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	authInputForm := &user.User{
		Login:    login,
		Username: username,
		Password: password,
	}
	h.serviceUser.Create(authInputForm)

	if login == "" || password == "" || username == "" {
		c.HTML(http.StatusBadRequest, "registration", gin.H{"ErrorMessage": "Parameters can't be empty"})
		return
	}
	c.Redirect(http.StatusFound, "/")
}
