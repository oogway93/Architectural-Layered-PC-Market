package repositoryPostgresUser

import (
	"log"

	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
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

	result := d.db.Where("login = ? AND deleted_at IS NULL", loginID).First(&user)
	if result.Error != nil {
		log.Printf("Error finding CERTAIN LOGIN from user table: %v", result.Error)
	}

	resultUser := map[string]interface{}{
		"login":         user.Login,
		"hash_password": user.Password,
	}
	tx.Commit()
	return resultUser
}
