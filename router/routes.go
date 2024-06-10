package router

import (
	"github.com/RianIhsan/go-topup-midtrans/feature/auth"
	"github.com/RianIhsan/go-topup-midtrans/feature/users"
	"github.com/RianIhsan/go-topup-midtrans/middleware"
	"github.com/RianIhsan/go-topup-midtrans/utils/jwtToken"
	"github.com/gofiber/fiber/v2"
)

func BootAuthRoute(app *fiber.App, handler auth.AuthHandlerInterface) {
	authGroup := app.Group("/v1/auth")
	authGroup.Post("/register", handler.Register)
	authGroup.Post("/login", handler.Login)
}

func BootUserRoute(app *fiber.App, handler users.UserHandlerInterface, jwtMiddleware jwtToken.IJwt, userService users.UserServiceInterface) {
	userGroup := app.Group("/v1/user")
	userGroup.Get("/me", middleware.Protected(jwtMiddleware, userService), handler.GetCurrentUser)
}
