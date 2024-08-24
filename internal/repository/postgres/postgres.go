package repositoryPostgres

import (
	"fmt"
	"log"

	"github.com/oogway93/golangArchitecture/internal/repository"
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/shop"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

func DatabaseConnection(cfg Config) (*gorm.DB, error) {

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error with initialized gorm db: %v", err)
		return nil, err
	}

	return db, nil
}

func NewRepository(db *gorm.DB) *repository.Repository {
	return &repository.Repository{
		CategoryRepository: repositoryPostgresShop.NewRepositoryCategoryShop(db),
	}
}
