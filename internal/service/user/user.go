package serviceUser

import (
	"github.com/oogway93/golangArchitecture/internal/entity/user"
	"github.com/oogway93/golangArchitecture/internal/repository"
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
)

type UserService struct {
	repositoryUser repository.UserRepository
}

func NewServiceUser(repo repository.UserRepository) *UserService {
	return &UserService{
		repositoryUser: repo,
	}
}

func (c *UserService) Create(requestData *user.User) {
	userModel := models.User{
		Login:    requestData.Login,
		Username: requestData.Username,
		Password: requestData.Password,
	}
	c.repositoryUser.Create(userModel)
}

func (c *UserService) GetAll() []map[string]interface{} {
	result := c.repositoryUser.GetAll()
	return result
}

func (c *UserService) Update(loginId string, requestData *user.UserUpdated) error {
	userModel := models.User{
		Username: requestData.Username,
		Password: requestData.Password,
	}
	status := c.repositoryUser.Update(loginId, userModel)
	return status
}

func (c *UserService) Delete(loginID string) error {
	result := c.repositoryUser.Delete(loginID)
	return result
}
