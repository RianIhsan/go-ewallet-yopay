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

func (a authHandler) Login(c *fiber.Ctx) error {
	var payload dto.LoginRequest

	if err := c.BodyParser(&payload); err != nil {
		return response.SendStatusBadRequest(c, "invalid payload:"+err.Error())
	}
	if err := validator.ValidateStruct(payload); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}

	userLogin, token, err := a.authService.Login(&payload)
	if err != nil {
		if err.Error() == "user not found" {
			return response.SendStatusNotFound(c, "user not found")
		}
		return response.SendStatusUnauthorized(c, "incorrect password")
	}
	return response.SendStatusOkWithDataResponse(c, "login is successfully", dto.LoginResponse(userLogin, token))

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
