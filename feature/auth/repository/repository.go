package repository

import (
	"github.com/RianIhsan/go-topup-midtrans/entities"
	"github.com/RianIhsan/go-topup-midtrans/feature/auth"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func (a authRepository) InsertUser(newUser *entities.MstUser) (*entities.MstUser, error) {
	if err := a.db.Create(newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil
}

func NewAuthRepository(db *gorm.DB) auth.AuthRepositoryInterface {
	return &authRepository{db}
}
