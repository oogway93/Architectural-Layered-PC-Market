package serviceShop

import (
	"fmt"
	"log/slog"

	"github.com/oogway93/golangArchitecture/internal/core/entity/products"
	"github.com/oogway93/golangArchitecture/internal/core/repository"
	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
	"github.com/oogway93/golangArchitecture/internal/core/utils"
)

type ProductShopService struct {
	repo  repository.ProductRepository
	cache repository.CacheRepository
}

func NewServiceShopProduct(repo repository.ProductRepository, cache repository.CacheRepository) *ProductShopService {
	return &ProductShopService{
		repo:  repo,
		cache: cache,
	}
}

func (s *ProductShopService) Create(categoryID string, requestData *products.Product) {
	productModel := models.Product{
		ProductName: requestData.ProductName,
		Price:       requestData.Price,
		Description: requestData.Description,
	}
	s.repo.Create(categoryID, &productModel)
}
func (s *ProductShopService) GetAll(reqFrom string, categoryID string) ([]models.Product, []map[string]interface{}) {
	if reqFrom == "HTTP" {
		var productsModel []models.Product
		key := "productsHTTP"
		cachedProducts, err := s.cache.Get(key)
		if err == nil {
			err := utils.Deserialize(cachedProducts, &productsModel)
			if err != nil {
				return nil, nil
			}
			return productsModel, nil
		}
		productsModel, _ = s.repo.GetAll(categoryID)
		if len(productsModel) != 0 {
			productsSerialized, err := utils.Serialize(productsModel)
			if err != nil {
				slog.Warn("serialization incorrect")
			}
			err = s.cache.Set(key, productsSerialized)
			if err != nil {
				slog.Warn("set cache incorrect")
			}
		}
		return productsModel, nil
	}
	var products []map[string]interface{}
	key := "productsAPI"
	cachedProducts, err := s.cache.Get(key)
	if err == nil {
		err := utils.Deserialize(cachedProducts, &products)
		if err != nil {
			return nil, nil
		}
		return nil, products
	}
	_, products = s.repo.GetAll(categoryID)
	if len(products) != 0 {
		productsSerialized, err := utils.Serialize(products)
		if err != nil {
			slog.Warn("serialization incorrect")
		}
		err = s.cache.Set(key, productsSerialized)
		if err != nil {
			slog.Warn("set cache incorrect")
		}
	}
	return nil, products
}
func (s *ProductShopService) Delete(categoryID, productID string) error {
	key := fmt.Sprintf("category:%s::product:%s", categoryID, productID)
	err := s.cache.Delete(key)
	if err != nil {
		return fmt.Errorf("error in Delete  method category cache")
	}
	err = s.repo.Delete(categoryID, productID)
	if err != nil {
		return fmt.Errorf("error in Delete  method category repo postgres")
	}

	return err
}
func (s *ProductShopService) Get(categoryID, productID string) map[string]interface{} {
	var product map[string]interface{}
	key := fmt.Sprintf("category:%s::product:%s", categoryID, productID)
	cachedProduct, err := s.cache.Get(key)
	if err == nil {
		err := utils.Deserialize(cachedProduct, &product)
		if err != nil {
			return nil
		}

		return product
	}
	product = s.repo.Get(categoryID, productID)

	productSerialized, err := utils.Serialize(product)
	if err != nil {
		return nil
	}

	err = s.cache.Set(key, productSerialized)
	if err != nil {
		return nil
	}
	return product
}
func (s *ProductShopService) Update(categoryID, productID string, requestData *products.Product) error {
	productModel := models.Product{
		ProductName: requestData.ProductName,
		Price:       requestData.Price,
		Description: requestData.Description,
	}
	newCategoryName := requestData.CategoryRel.CategoryName
	resultProduct, err := s.repo.Update(newCategoryName, productID, productModel)
	if err != nil {
		return fmt.Errorf("error in Update method category repo")
	}
	key := fmt.Sprintf("category:%s::product:%s", categoryID, productID)
	err = s.cache.Delete(key)
	if err != nil {
		return fmt.Errorf("error in Update  method category cache")
	}

	if requestData.CategoryRel.CategoryName == "" {
		categoryName := s.repo.GetByCategoryId(resultProduct["categoryID"].(uint))
		resultProduct["category_name"] = categoryName
	}

	productSerialized, err := utils.Serialize(resultProduct)
	if err != nil {
		return fmt.Errorf("error in Serialization Update method category cache")
	}
	newKey := fmt.Sprintf("category:%s::product:%s", resultProduct["category_name"], resultProduct["product_name"])
	err = s.cache.Set(newKey, productSerialized)
	if err != nil {
		slog.Warn("set cache incorrect")
	}
	return nil
}
