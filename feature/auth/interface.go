package auth

import (
	"github.com/RianIhsan/go-topup-midtrans/entities"
	"github.com/RianIhsan/go-topup-midtrans/feature/auth/dto"
	"github.com/gofiber/fiber/v2"
)

type (
	AuthRepositoryInterface interface {
		InsertUser(newUser *entities.MstUser) (*entities.MstUser, error)
	}
	AuthServiceInterface interface {
		Register(newUser *dto.RegisterRequest) (*entities.MstUser, error)
	}
	AuthHandlerInterface interface {
		Register(c *fiber.Ctx) error
	}
)
