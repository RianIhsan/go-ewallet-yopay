package dto

import "github.com/RianIhsan/go-topup-midtrans/entities"

type TypeLoginResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func LoginResponse(user *entities.MstUser, token string) *TypeLoginResponse {
	return &TypeLoginResponse{
		Fullname: user.Fullname,
		Email:    user.Email,
		Token:    token,
	}
}
