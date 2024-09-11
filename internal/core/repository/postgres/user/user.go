package repositoryPostgresUser

import (
	"log"
	"time"

	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
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

func (d *UserPostgres) Create(newUser models.User) {
	tx := d.db.Begin()

	result := tx.Create(&newUser)

	if result.Error != nil {
		log.Printf("Error creating new user: %v", result.Error)
	}
	tx.Commit()
}

func (d *UserPostgres) GetAll() []map[string]interface{} {
	var users []models.User
	tx := d.db.Begin()
	result := tx.Find(&users)

	if result.Error != nil {
		log.Printf("Error finding records from user: %v", result.Error)
	}
	var resultUsers []map[string]interface{}
	for _, user := range users {
		resultUsers = append(resultUsers, map[string]interface{}{
			"username": user.Username,
			"password": user.Password})
	}
	tx.Commit()
	return resultUsers
}

func (d *UserPostgres) Get(loginID string) map[string]interface{} {
	var user models.User
	tx := d.db.Begin()
	result := tx.Where("login = ?", loginID).First(&user)
	if result.Error != nil {
		log.Printf("Error finding LOGIN from user: %v", result.Error)
		return nil
	}

	resultUser := map[string]interface{}{
		"id":    user.ID,
		"login": user.Login,
	}

	tx.Commit()
	return resultUser
}

func (d *UserPostgres) Update(loginID string, newUser models.User) error {
	var user models.User
	tx := d.db.Begin()
	result := tx.Where("login = ?", loginID).First(&user)
	if result.Error != nil {
		return result.Error
	}
	user.Username = newUser.Username
	user.Password = newUser.Password
	user.UpdatedAt = time.Now()

	result = tx.Save(&user)
	tx.Commit()
	return result.Error
}

func (d *UserPostgres) Delete(loginID string) error {
	var user models.User
	tx := d.db.Begin()
	result := tx.Where("login = ?", loginID).Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	tx.Commit()
	return result.Error
}
