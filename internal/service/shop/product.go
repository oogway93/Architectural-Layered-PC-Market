package serviceShop

import "github.com/oogway93/golangArchitecture/internal/repository/postgres/models"

type CategoryShopService struct {
	repositoryShopCategory repository.CategoryRepository
}

func NewServiceShopCategory(repo repository.Pro) *CategoryShopService {
	return &CategoryShopService{
		repositoryShopCategory: repo,
	}
}

func (c *CategoryShopService) Create(requestData models.Category) 
func (c *CategoryShopService) GetAll() []map[string]interface{}
func (c *CategoryShopService) Delete(categoryID string) error
func (c *CategoryShopService) Get(categoryID string) string
func (c *CategoryShopService) Update(categoryID string, newCategory models.Category) error
