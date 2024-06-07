package entities

import "gorm.io/gorm"

type MstUser struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	gorm.Model
}

func (MstUser) TableName() string {
	return "mst_user"
}
