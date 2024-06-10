package dto

import "github.com/RianIhsan/go-topup-midtrans/entities"

type UserBalanceResponse struct {
	Id           int     `json:"id"`
	Fullname     string  `json:"fullname"`
	Email        string  `json:"email"`
	Phone        string  `json:"phone"`
	TotalBalance float64 `json:"total_balance"`
}

func GetTotalBalanceUser(user *entities.MstUser, totalBalance float64) *UserBalanceResponse {
	return &UserBalanceResponse{
		Id:           user.Id,
		Fullname:     user.Fullname,
		Email:        user.Email,
		Phone:        user.Phone,
		TotalBalance: totalBalance,
	}
}
