package serviceShop

import (
	"log"
	"time"

	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/repository"
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
	"github.com/oogway93/golangArchitecture/internal/utils"
)

const (
	ttl = 5 * time.Minute
)

type CategoryShopService struct {
	repo  repository.CategoryRepository
	cache repository.CacheRepository
}

func NewServiceShopCategory(repo repository.CategoryRepository, cache repository.CacheRepository) *CategoryShopService {
	return &CategoryShopService{
		repo:  repo,
		cache: cache,
	}
}

func (s *CategoryShopService) Create(requestData *products.Category) {
	categoryModel := models.Category{
		CategoryName: requestData.CategoryName,
	}
	s.repo.Create(categoryModel)
	key := "category"
	categorySerialized, err := utils.Serialize(categoryModel)
	if err != nil {
		log.Fatal("serialization incorrect")
	}
	err = s.cache.Set(key, categorySerialized, ttl)
	if err != nil {
		log.Fatal("set cache incorrect")
	}
}

func (s *CategoryShopService) GetAll() []map[string]interface{} {
	var categories []map[string]interface{}
	key := "categories"
	cachedCategories, err := s.cache.Get(key)
	if err == nil {
		err := utils.Deserialize(cachedCategories, &categories)
		if err != nil {
			return nil
		}

		return categories
	}
	categories = s.repo.GetAll()

	categoriesSerialized, err := utils.Serialize(categories)
	if err != nil {
		return nil
	}

	err = s.cache.Set(key, categoriesSerialized, ttl)
	if err != nil {
		return nil
	}

	return categories
}
func (s *CategoryShopService) Get(categoryID string) map[string]interface{} {
	result := s.repo.Get(categoryID)
	return result
}

func (s *CategoryShopService) Delete(categoryID string) error {
	result := s.repo.Delete(categoryID)
	return result
}
func (s *CategoryShopService) Update(categoryID string, requestData *products.Category) error {
	categoryModel := models.Category{
		CategoryName: requestData.CategoryName,
	}
	result := s.repo.Update(categoryID, categoryModel)
	return result
}
