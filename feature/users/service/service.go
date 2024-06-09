package service

import (
	"errors"
	"github.com/RianIhsan/go-topup-midtrans/entities"
	"github.com/RianIhsan/go-topup-midtrans/feature/users"
)

type userService struct {
	repo users.UserRepositoryInterface
}

func (u userService) GetId(id int) (*entities.MstUser, error) {
	result, err := u.repo.FindId(id)
	if err != nil {
		return nil, errors.New("id not found")
	}
	return result, nil
}

func (u userService) GetEmail(email string) (*entities.MstUser, error) {
	result, err := u.repo.FindEmail(email)
	if err != nil {
		return nil, errors.New("your email has been already")
	}
	return result, nil
}

func NewUserService(repo users.UserRepositoryInterface) users.UserServiceInterface {
	return &userService{repo}
}
