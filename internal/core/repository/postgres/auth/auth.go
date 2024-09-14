package repositoryPostgresAuth

import (
	"log/slog"

	// "github.com/oogway93/golangArchitecture/internal/repository"
	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewRepositoryAuth(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (d *AuthPostgres) Login(loginID string) map[string]interface{} {
	var user models.User
	tx := d.db.Begin()

	result := tx.Where("login = ? AND deleted_at IS NULL", loginID).First(&user)
	if result.Error != nil {
		slog.Warn("Error finding CERTAIN LOGIN from user table", "error", result.Error)
	}

	resultUser := map[string]interface{}{
		"login":         user.Login,
		"hash_password": user.Password,
	}
	tx.Commit()
	return resultUser
}
