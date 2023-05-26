package usecase

import (
	"fmt"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/repository"
)

type UserUseCase interface {
	RegisterNewUser(newUser entity.UserCredential) error
	FindAllUser() ([]entity.UserCredential, error)
	GetUser(id string) (entity.UserCredential, error)
	UpdateUser(newUser entity.UserCredential) error
	DeleteUser(id string) error
	FindUserByUsername(username string) (entity.UserCredential, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func (u *userUseCase) RegisterNewUser(newUser entity.UserCredential) error {
	isUsernameExist, _ := u.repo.GetByUsername(newUser.UserName)
	if isUsernameExist.UserName == newUser.UserName {
		return fmt.Errorf("user with useername: %v exists", newUser.UserName)
	}
	if newUser.UserName == "" || newUser.Password == "" {
		return fmt.Errorf("userName and Password are required fields")
	}
	// create user credential (recommended use transactional)
	//password, err := security.HashPassword(newUser.Password)
	//if err != nil {
	//	return err
	//}
	newUser.IsActive = true
	//newUser.Password = password
	err := u.repo.Create(newUser)
	if err != nil {
		return fmt.Errorf("failed to create new user: %v", err)
	}

	return nil
}

func (u *userUseCase) FindAllUser() ([]entity.UserCredential, error) {
	return u.repo.List()
}

func (u *userUseCase) GetUser(id string) (entity.UserCredential, error) {
	return u.repo.Get(id)
}

func (u *userUseCase) UpdateUser(newUser entity.UserCredential) error {
	isUsernameExist, _ := u.repo.GetByUsername(newUser.UserName)
	if isUsernameExist.UserName == newUser.UserName && isUsernameExist.Id != newUser.Id {
		return fmt.Errorf("customer with username: %v exists", newUser.UserName)
	}
	if newUser.UserName == "" || newUser.Password == "" {
		return fmt.Errorf("UserName and Password are required fields")
	}
	err := u.repo.Update(newUser)
	if err != nil {
		return fmt.Errorf("failed to udpate vehicle: %v", err)
	}

	return nil
}

func (u *userUseCase) DeleteUser(id string) error {
	return u.repo.Delete(id)
}

func (u *userUseCase) FindUserByUsername(username string) (entity.UserCredential, error) {
	user, err := u.repo.GetByUsername(username)
	if err != nil {
		return entity.UserCredential{}, fmt.Errorf("user with username %s not found", username)
	}
	return user, nil
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
