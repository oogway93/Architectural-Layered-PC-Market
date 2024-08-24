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

func  (c *UserService) Create(requestData *user.User) {
	userModel := models.User{
		ID: 1,
		Username: requestData.Username,
		Password: requestData.Password,
	}
	c.repositoryUser.Create(userModel)

}
