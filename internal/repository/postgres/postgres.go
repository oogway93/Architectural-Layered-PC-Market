package repositoryPostgres

import (
	"fmt"
	"log"

	"github.com/oogway93/golangArchitecture/internal/repository"
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/shop"
	repositoryPostgresUser "github.com/oogway93/golangArchitecture/internal/repository/postgres/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
	SSLMode  string
}

func DatabaseConnection(cfg Config) *gorm.DB {

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
	return db
}

func NewRepository(db *gorm.DB) *repository.Repository {
	return &repository.Repository{
		CategoryRepository: repositoryPostgresShop.NewRepositoryCategoryShop(db),
		UserRepository: repositoryPostgresUser.NewRepositoryUser(db),
	}
}
