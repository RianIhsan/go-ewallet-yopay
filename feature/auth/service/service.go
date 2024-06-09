package service

import (
	"errors"
	"github.com/RianIhsan/go-topup-midtrans/entities"
	"github.com/RianIhsan/go-topup-midtrans/feature/auth"
	"github.com/RianIhsan/go-topup-midtrans/feature/auth/dto"
	"github.com/RianIhsan/go-topup-midtrans/feature/users"
	"github.com/RianIhsan/go-topup-midtrans/utils/hashing"
)

type authService struct {
	repo        auth.AuthRepositoryInterface
	userService users.UserServiceInterface
	hashing     hashing.HashInterface
}

func (a authService) Register(newUser *dto.RegisterRequest) (*entities.MstUser, error) {
	isExistEmail, _ := a.userService.GetEmail(newUser.Email)
	if isExistEmail != nil {
		return nil, errors.New("email already exist")
	}
	hashedPassword, err := a.hashing.GenerateHash(newUser.Password)
	if err != nil {
		return nil, err
	}
	data := &entities.MstUser{
		Fullname: newUser.Fullname,
		Email:    newUser.Email,
		Password: hashedPassword,
		Phone:    newUser.Phone,
	}
	user, err := a.repo.InsertUser(data)
	if err != nil {
		return nil, errors.New("failed create account")
	}
	return user, nil
}

func NewAuthService(
	repo auth.AuthRepositoryInterface,
	userService users.UserServiceInterface,
	hashing hashing.HashInterface) auth.AuthServiceInterface {
	return &authService{repo, userService, hashing}
}
