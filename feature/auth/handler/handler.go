package handler

import (
	"github.com/RianIhsan/go-topup-midtrans/feature/auth"
	"github.com/RianIhsan/go-topup-midtrans/feature/auth/dto"
	"github.com/RianIhsan/go-topup-midtrans/utils/response"
	"github.com/RianIhsan/go-topup-midtrans/utils/validator"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService auth.AuthServiceInterface
}

func (a authHandler) Register(c *fiber.Ctx) error {
	var payload dto.RegisterRequest

	if err := c.BodyParser(&payload); err != nil {
		return response.SendStatusBadRequest(c, "invalid payload:"+err.Error())
	}

	if err := validator.ValidateStruct(payload); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}

	_, err := a.authService.Register(&payload)
	if err != nil {
		return response.SendStatusBadRequest(c, err.Error())
	}
	return response.SendStatusOkResponse(c, "register is successfully")
}

func NewAuthHandler(authService auth.AuthServiceInterface) auth.AuthHandlerInterface {
	return &authHandler{authService}
}
