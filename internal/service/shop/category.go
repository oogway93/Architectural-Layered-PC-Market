package serviceShop

import (
	"fmt"
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
	var category map[string]interface{}
	key := fmt.Sprintf("category:%s", categoryID)
	cachedCategories, err := s.cache.Get(key)
	if err == nil {
		err := utils.Deserialize(cachedCategories, &category)
		if err != nil {
			return nil
		}

		return category
	}
	category = s.repo.Get(categoryID)

	categoriesSerialized, err := utils.Serialize(category)
	if err != nil {
		return nil
	}

	err = s.cache.Set(key, categoriesSerialized, ttl)
	if err != nil {
		return nil
	}
	return category
}
func (s *CategoryShopService) Delete(categoryID string) error {
	key := fmt.Sprintf("category:%s", categoryID)
	err := s.cache.Delete(key)
	if err != nil {
		return fmt.Errorf("error in Delete  method category cache")
	}
	err = s.repo.Delete(categoryID)

	if err != nil {
		return fmt.Errorf("error in Delete  method category repo postgres")	}

	return nil
}
func (s *CategoryShopService) Update(categoryID string, requestData *products.Category) error {
	categoryModel := models.Category{
		CategoryName: requestData.CategoryName,
	}
	
	err := s.repo.Update(categoryID, categoryModel)
	if err != nil {
		return fmt.Errorf("error in Update method category repo")
	}

	key := fmt.Sprintf("category:%s", categoryID)
	err = s.cache.Delete(key)
	if err != nil {
		return fmt.Errorf("error in Update  method category cache")
	}

	categorySerialized, err := utils.Serialize(requestData.CategoryName)
	if err != nil {
		return fmt.Errorf("error in Serilization Update method category cache")
	}

	newKey := fmt.Sprintf("category:%s", requestData.CategoryName)
	err = s.cache.Set(newKey, categorySerialized, ttl)
	if err != nil {
		log.Fatal("set cache incorrect")
	}
	return nil

}
