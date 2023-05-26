package usecase

import (
	"fmt"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/repository"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/utils/security"
)

type AuthenticationUseCase interface {
	Login(username string, password string) (string, error)
}

type authenticationUseCase struct {
	repo         repository.UserRepository
	tokenService security.AccessToken
}

func (a *authenticationUseCase) Login(username string, password string) (string, error) {
	user, err := a.repo.GetByUsernamePassword(username, password)
	if err != nil {
		return "", fmt.Errorf("user with username: %s not found", username)
	}
	var token string
	token, err = a.tokenService.CreateAccessToken(&user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewAuthenticationUseCase(repo repository.UserRepository, tokenService security.AccessToken) AuthenticationUseCase {
	return &authenticationUseCase{repo: repo, tokenService: tokenService}
}
