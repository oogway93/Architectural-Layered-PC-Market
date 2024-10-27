package HTTPShopCategoryHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oogway93/golangArchitecture/internal/core/entity/HTTP/shop"
	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
	// "github.com/oogway93/golangArchitecture/internal/core/server/serverAPI/handler/auth"
)

type TemplateData struct {
	Categories       []models.Category
	CategoryProducts productsHTTP.CategoryProducts
	Token            string
}

func (h *HTTPCategoryHandler) GetAll(c *gin.Context) {
	resultCategoryModel, _ := h.serviceCategory.GetAll("HTTP")
	templateData := &TemplateData{Categories: resultCategoryModel}
	c.HTML(http.StatusOK, "categories", templateData)
}

func (h *HTTPCategoryHandler) Get(c *gin.Context) {
	categoryID := c.Param("category")
	resultCategory, _ := h.serviceCategory.Get("HTTP", categoryID)
	resultProductsModels, _ := h.serviceProduct.GetAll("HTTP", categoryID)

	categoryProducts := productsHTTP.CategoryProducts{
		Category: resultCategory,
		Products: resultProductsModels,
	}
	// token, err := handlerAuth.GetJWTToken(c)
	// if err != nil {
	// 	c.Redirect(http.StatusSeeOther, "/login")
	// 	return
	// }

	// Validate the token here
	// claims, err := handlerAuth.validateJWT(token)
	// if err != nil {
	//     c.Redirect(http.StatusSeeOther, "/login")
	//     return
	// }

	// // Render the dashboard template with user data
	// c.HTML(http.StatusOK, "dashboard.html", gin.H{
	//     "username": claims["username"].(string),
	// })
	templateData := &TemplateData{CategoryProducts: categoryProducts}
	c.HTML(http.StatusOK, "categoryWithProducts", templateData)
}
