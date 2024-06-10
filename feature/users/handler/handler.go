package handler

import (
	"github.com/RianIhsan/go-topup-midtrans/entities"
	"github.com/RianIhsan/go-topup-midtrans/feature/users"
	"github.com/RianIhsan/go-topup-midtrans/feature/users/dto"
	"github.com/RianIhsan/go-topup-midtrans/utils/response"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService users.UserServiceInterface
}

func (u userHandler) GetCurrentUser(c *fiber.Ctx) error {
	user, ok := c.Locals("CurrentUser").(*entities.MstUser)
	if !ok || user == nil {
		return response.SendStatusUnauthorized(c, "Access denied: user not found")
	}
	return response.SendStatusOkWithDataResponse(c, "current user", dto.GetUserResponse(user))
}

func NewUserHandler(userService users.UserServiceInterface) users.UserHandlerInterface {
	return &userHandler{
		userService: userService,
	}
}
