package entities

import "gorm.io/gorm"

type MstWithdrawBalance struct {
	Id           int     `gorm:"primaryKey" json:"id"`
	UserId       int     `gorm:"not null" json:"user_id"`
	Token        int     `gorm:"not null" json:"token"`
	Amount       float64 `gorm:"not null" json:"amount"`
	Provider     string  `gorm:"not null" json:"provider"`
	Status       string  `gorm:"type:varchar(50);default:'pending'" json:"status"`
	TokenExpired int64   `gorm:"column:token_expired;type:bigint" json:"token_expired" `
	gorm.Model
}

func (MstWithdrawBalance) TableName() string {
	return "mst_withdraw_balance"
}
