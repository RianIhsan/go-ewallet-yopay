package service

import (
	"errors"
	"github.com/RianIhsan/go-topup-midtrans/entities"
	"github.com/RianIhsan/go-topup-midtrans/feature/auth"
	"github.com/RianIhsan/go-topup-midtrans/feature/auth/dto"
	"github.com/RianIhsan/go-topup-midtrans/feature/users"
	"github.com/RianIhsan/go-topup-midtrans/utils/hashing"
	"github.com/RianIhsan/go-topup-midtrans/utils/jwtToken"
)

type authService struct {
	repo        auth.AuthRepositoryInterface
	userService users.UserServiceInterface
	hashing     hashing.HashInterface
	jwt         jwtToken.IJwt
}

func (a authService) Login(user *dto.LoginRequest) (*entities.MstUser, string, error) {
	userData, err := a.userService.GetEmail(user.Email)
	if err != nil {
		return nil, "", errors.New("email not found")
	}

	isValidPassword, err := a.hashing.ComparePassword(userData.Password, user.Password)
	if err != nil || !isValidPassword {
		return nil, "", errors.New("incorrect password")
	}

	accessSecret, err := a.jwt.GenerateJWT(userData.Id, userData.Email, userData.Fullname)
	if err != nil {
		return nil, "", err
	}
	return userData, accessSecret, nil
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
	hashing hashing.HashInterface,
	jwtToken jwtToken.IJwt) auth.AuthServiceInterface {
	return &authService{repo, userService, hashing, jwtToken}
}
