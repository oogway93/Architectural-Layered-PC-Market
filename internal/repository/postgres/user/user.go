package repositoryPostgresUser

import (
	"log"

	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (d *UserPostgres) Create(user models.User) {
	result := d.db.Create(&user)
	
	if result.Error != nil {
        log.Printf("Error creating user: %v", result.Error)
}


}