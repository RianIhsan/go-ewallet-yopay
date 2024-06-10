package dto

import "github.com/RianIhsan/go-topup-midtrans/entities"

type TypeGetUserResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func GetUserResponse(user *entities.MstUser) *TypeGetUserResponse {
	return &TypeGetUserResponse{
		Fullname: user.Fullname,
		Email:    user.Email,
		Phone:    user.Phone,
	}
}
