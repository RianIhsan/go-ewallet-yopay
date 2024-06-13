package router

import (
	"github.com/RianIhsan/go-topup-midtrans/feature/auth"
	"github.com/RianIhsan/go-topup-midtrans/feature/topup"
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

func BootTopUpBalanceRoute(app *fiber.App, handler topup.TopUpHandlerInterface, jwtMiddleware jwtToken.IJwt, userService users.UserServiceInterface) {
	topUpGroup := app.Group("/v1/balance")
	topUpGroup.Post("/add", middleware.Protected(jwtMiddleware, userService), handler.CreateTopUp)
	topUpGroup.Post("/callback", handler.CallBack)
	topUpGroup.Get("/total", middleware.Protected(jwtMiddleware, userService), handler.GetTotalBalance)
	topUpGroup.Post("/transfer", middleware.Protected(jwtMiddleware, userService), handler.TransferBalance)
	topUpGroup.Post("/withdraw", middleware.Protected(jwtMiddleware, userService), handler.CreateTokenWithdraw)
	topUpGroup.Post("/confirm-withdraw", middleware.Protected(jwtMiddleware, userService), handler.ConfirmWithdraw)
}
