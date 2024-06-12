package repository

import (
	"github.com/RianIhsan/go-topup-midtrans/entities"
	"github.com/RianIhsan/go-topup-midtrans/feature/users"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (r userRepository) FindEmail(email string) (*entities.MstUser, error) {
	var user *entities.MstUser
	if err := r.db.Table("mst_user").
		Where("email = ? AND deleted_at IS NULL", email).
		First(&user).
		Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r userRepository) FindId(id int) (*entities.MstUser, error) {
	var user *entities.MstUser
	if err := r.db.
		Where("id = ? AND deleted_at IS NULL", id).
		First(&user).
		Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r userRepository) FindUserByPhone(phone string) (*entities.MstUser, error) {
	var user *entities.MstUser
	if err := r.db.Table("mst_user").
		Where("phone = ? AND deleted_at IS NULL", phone).
		First(&user).
		Error; err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) users.UserRepositoryInterface {
	return &userRepository{db}
}
