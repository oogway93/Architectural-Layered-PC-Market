package repositoryPostgresUser

import (
	"log"
	"time"

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

func (d *UserPostgres) Create(newUser models.User) {
	tx := d.db.Begin()

	result := d.db.Create(&newUser)	

	if result.Error != nil {
		log.Printf("Error creating new user: %v", result.Error)
	}
	tx.Commit()
}

func (d *UserPostgres) GetAll() []map[string]interface{} {
	var users []models.User
	tx := d.db.Begin()
	result := d.db.Find(&users)

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

func (d *UserPostgres) Update(loginID string, newUser models.User) (error){
	var user models.User
	tx := d.db.Begin()
	result := d.db.Where("login = ?", loginID).First(&user)
    if result.Error != nil {
        // Handle error (e.g., user not found)
        return result.Error
    }
	user.Username = newUser.Username
    user.Password = newUser.Password
	user.UpdatedAt = time.Now()
    // Add any other fields you need to update

    result = d.db.Save(&user)
	tx.Commit()
    return result.Error
}

func (d *UserPostgres) Delete(loginID string) error {
	var user models.User
	tx := d.db.Begin()
	result := d.db.Where("login = ?", loginID).Delete(&user)
	if result.Error != nil {
        // Handle error (e.g., user not found)
        return result.Error
    }
	tx.Commit()
	return result.Error
}
