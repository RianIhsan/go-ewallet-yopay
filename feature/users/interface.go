package users

import "github.com/RianIhsan/go-topup-midtrans/entities"

type (
	UserRepositoryInterface interface {
		FindEmail(email string) (*entities.MstUser, error)
		FindId(id int) (*entities.MstUser, error)
	}
	UserServiceInterface interface {
		GetId(id int) (*entities.MstUser, error)
		GetEmail(email string) (*entities.MstUser, error)
	}
)
