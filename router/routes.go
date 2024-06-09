package router

import (
	"github.com/RianIhsan/go-topup-midtrans/feature/auth"
	"github.com/gofiber/fiber/v2"
)

func BootAuthRoute(app *fiber.App, handler auth.AuthHandlerInterface) {
	authGroup := app.Group("/v1/auth")
	authGroup.Post("/register", handler.Register)
	//authGroup.Post("/login", handler.Login)
}
