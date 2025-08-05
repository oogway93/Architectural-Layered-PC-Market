package repositoryPostgresShop_test

import (
	"testing"

	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
	repo "github.com/oogway93/golangArchitecture/internal/core/repository/postgres/shop"
	"github.com/oogway93/golangArchitecture/internal/testdb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type CategoryRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo *repo.CategoryShopPostgres
}

func (s *CategoryRepositoryTestSuite) SetupTest() {
	s.db = testdb.WithTestDB(s.T())
	s.repo = repo.NewRepositoryCategoryShop(s.db)

	// Очищаем таблицу перед каждым тестом
	s.db.Exec("TRUNCATE TABLE categories RESTART IDENTITY CASCADE")
}

func (s *CategoryRepositoryTestSuite) TestCreateMethod() {
	category := &models.Category{CategoryName: "TEST_CATEGORY"}
	err := s.repo.Create(category)

	assert.NoError(s.T(), err)
	assert.NotZero(s.T(), category.ID)

	// Проверяем, что запись действительно создана
	var count int64
	s.db.Model(&models.Category{}).Count(&count)
	assert.Equal(s.T(), int64(1), count)
}

func (s *CategoryRepositoryTestSuite) TestFindByCategoryName() {
	// Создаем тестовые данные
	category := &models.Category{CategoryName: "TEST_CATEGORY"}
	err := s.repo.Create(category)
	assert.NoError(s.T(), err)
	// Ищем категорию
	found, _, err := s.repo.Get(category.CategoryName)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), int(category.ID), int(found.ID))
	assert.Equal(s.T(), category.CategoryName, found.CategoryName)
	assert.Equal(s.T(), len(found.Products), 0)
}

func (s *CategoryRepositoryTestSuite) TestFindAllRecords() {
	// Создаем тестовые данные
	category1 := &models.Category{CategoryName: "TEST_CATEGORY"}
	err := s.repo.Create(category1)
	assert.NoError(s.T(), err)
	category2 := &models.Category{CategoryName: "TEST_CATEGORY2"}
	err = s.repo.Create(category2)
	assert.NoError(s.T(), err)

	founds, jsonFounds, err := s.repo.GetAll()

	assert.NotNil(s.T(), jsonFounds)
	assert.Equal(s.T(), len(jsonFounds), 2)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), len(founds), 2)

	assert.Equal(s.T(), founds[0].CategoryName, category1.CategoryName)
	assert.True(s.T(), founds[0].DeletedAt.Time.IsZero())
	assert.Equal(s.T(), int(founds[0].ID), 1)
}

func (s *CategoryRepositoryTestSuite) TestDeleteMethod() {
	category := &models.Category{CategoryName: "TEST_CATEGORY"}
	err := s.repo.Create(category)
	assert.NoError(s.T(), err)

	err = s.repo.Delete("TEST_CATEGORY")
	assert.NoError(s.T(), err)
}
func TestCategoryRepositorySuite(t *testing.T) {
	suite.Run(t, new(CategoryRepositoryTestSuite))
}
