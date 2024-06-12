package users

import (
	"github.com/RianIhsan/go-topup-midtrans/entities"
	"github.com/gofiber/fiber/v2"
)

type (
	UserRepositoryInterface interface {
		FindEmail(email string) (*entities.MstUser, error)
		FindId(id int) (*entities.MstUser, error)
		FindUserByPhone(phone string) (*entities.MstUser, error)
	}
	UserServiceInterface interface {
		GetId(id int) (*entities.MstUser, error)
		GetEmail(email string) (*entities.MstUser, error)
		GetUserByPhone(phone string) (*entities.MstUser, error)
	}
	UserHandlerInterface interface {
		GetCurrentUser(c *fiber.Ctx) error
	}
)
