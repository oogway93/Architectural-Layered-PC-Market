package serviceShop

import (
	"fmt"
	"log/slog"

	"github.com/oogway93/golangArchitecture/internal/core/entity/API/shop"
	"github.com/oogway93/golangArchitecture/internal/core/repository"
	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
	"github.com/oogway93/golangArchitecture/internal/core/utils"
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

func (s *CategoryShopService) Create(requestData *productsAPI.Category) {
	categoryModel := models.Category{
		CategoryName: requestData.CategoryName,
	}
	s.repo.Create(&categoryModel)
}

func (s *CategoryShopService) GetAll(reqFrom string) ([]models.Category, []map[string]interface{}) {
	if reqFrom == "HTTP" {
		var categoriesModel []models.Category
		key := "categoriesHTTP"
		cachedCategories, err := s.cache.Get(key)
		if err == nil {
			err := utils.Deserialize(cachedCategories, &categoriesModel)
			if err != nil {
				return nil, nil
			}

			return categoriesModel, nil
		}
		categoriesModel, _, _ = s.repo.GetAll()

		if len(categoriesModel) != 0 {
			categoriesSerialized, err := utils.Serialize(categoriesModel)
			if err != nil {
				slog.Warn("serialization incorrect")
			}

			err = s.cache.Set(key, categoriesSerialized)
			if err != nil {
				slog.Warn("set cache incorrect")
			}

		}
		return categoriesModel, nil
	}
	var categories []map[string]interface{}
	key := "categoriesAPI"
	cachedCategories, err := s.cache.Get(key)
	if err == nil {
		err := utils.Deserialize(cachedCategories, &categories)
		if err != nil {
			return nil, nil
		}

		return nil, categories
	}
	_, categories, _ = s.repo.GetAll()

	if len(categories) != 0 {
		categoriesSerialized, err := utils.Serialize(categories)
		if err != nil {
			slog.Warn("serialization incorrect")
		}

		err = s.cache.Set(key, categoriesSerialized)
		if err != nil {
			slog.Warn("set cache incorrect")
		}

	}
	return nil, categories
}
func (s *CategoryShopService) Get(reqFrom, categoryID string) (models.Category, map[string]interface{}) {
	if reqFrom == "HTTP" {
		var categoryModel models.Category
		key := fmt.Sprintf("category:%s", categoryID)
		cachedCategories, err := s.cache.Get(key)
		if err == nil {
			_ = utils.Deserialize(cachedCategories, &categoryModel)
			return categoryModel, nil
		}
		categoryModel, _, _ = s.repo.Get(categoryID)

		categoriesSerialized, err := utils.Serialize(categoryModel)
		if err != nil {
			return models.Category{}, nil
		}

		err = s.cache.Set(key, categoriesSerialized)
		if err != nil {
			return models.Category{}, nil
		}
		return categoryModel, nil
	}
	var category map[string]interface{}
	key := fmt.Sprintf("category:%s", categoryID)
	cachedCategories, err := s.cache.Get(key)
	if err == nil {
		err := utils.Deserialize(cachedCategories, &category)
		if err != nil {
			return models.Category{}, nil
		}

		return models.Category{}, category
	}
	_, category, _ = s.repo.Get(categoryID)

	categoriesSerialized, err := utils.Serialize(category)
	if err != nil {
		return models.Category{}, nil
	}

	err = s.cache.Set(key, categoriesSerialized)
	if err != nil {
		return models.Category{}, nil
	}
	return models.Category{}, category
}
func (s *CategoryShopService) Delete(categoryID string) error {
	key := fmt.Sprintf("category:%s", categoryID)
	err := s.cache.Delete(key)
	if err != nil {
		return fmt.Errorf("error in Delete  method category cache")
	}
	err = s.repo.Delete(categoryID)

	if err != nil {
		return fmt.Errorf("error in Delete  method category repo postgres")
	}

	return nil
}
func (s *CategoryShopService) Update(categoryID string, requestData *productsAPI.Category) error {
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
	err = s.cache.Set(newKey, categorySerialized)
	if err != nil {
		slog.Warn("set cache incorrect", "error", err.Error)
	}
	return nil

}
