package serviceAuth

import (
	"github.com/oogway93/golangArchitecture/internal/entity/user"
	"github.com/oogway93/golangArchitecture/internal/repository"
	"github.com/oogway93/golangArchitecture/internal/utils"
)

type AuthService struct {
	repo  repository.AuthRepository
	cache repository.CacheRepository
}

func NewServiceAuth(repo repository.AuthRepository, cache repository.CacheRepository) *AuthService {
	return &AuthService{
		repo:  repo,
		cache: cache,
	}
}

func (s *AuthService) Login(requestData *user.AuthInput) bool {
	result := s.repo.Login(requestData.Login)
	checkValidationPassword := utils.CheckHashPassword(result["hash_password"].(string), requestData.Password)
	if checkValidationPassword {
		return true
	}
	return false
}
