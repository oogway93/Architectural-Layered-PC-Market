package serviceAuth

import (
	"github.com/oogway93/golangArchitecture/internal/entity/user"
	"github.com/oogway93/golangArchitecture/internal/repository"
	"github.com/oogway93/golangArchitecture/internal/utils"
)

type AuthService struct {
	repositoryAuth repository.AuthRepository
}

func NewServiceAuth(repo repository.AuthRepository) *AuthService {
	return &AuthService{
		repositoryAuth: repo,
	}
}

func (s *AuthService) Login(requestData *user.AuthInput) bool {
	result := s.repositoryAuth.Login(requestData.Login)
	checkValidationPassword := utils.CheckHashPassword(result["hash_password"].(string), requestData.Password)
	if checkValidationPassword {
		return true
	}
	return false
}
