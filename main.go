package main

import (
	"fmt"
	"github.com/RianIhsan/go-topup-midtrans/config"
	authHandler "github.com/RianIhsan/go-topup-midtrans/feature/auth/handler"
	authRepo "github.com/RianIhsan/go-topup-midtrans/feature/auth/repository"
	authService "github.com/RianIhsan/go-topup-midtrans/feature/auth/service"
	userRepo "github.com/RianIhsan/go-topup-midtrans/feature/users/repository"
	userService "github.com/RianIhsan/go-topup-midtrans/feature/users/service"
	"github.com/RianIhsan/go-topup-midtrans/router"
	"github.com/RianIhsan/go-topup-midtrans/utils/db"
	"github.com/RianIhsan/go-topup-midtrans/utils/hashing"
	"github.com/RianIhsan/go-topup-midtrans/utils/jwtToken"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:       "Welcome to API GOTOPUP",
		CaseSensitive: false,
	})
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Authorization",
		AllowOrigins:     "*",
		AllowCredentials: false,
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
	}))
	var initConfig = config.BootConfig()
	initDB := db.BootDatabase(*initConfig)
	db.MigrateTable(initDB)
	userRepository := userRepo.NewUserRepository(initDB)
	userService := userService.NewUserService(userRepository)
	hashing := hashing.NewHash()
	jwtInterface := jwtToken.NewJWT(initConfig.Secret)
	authRepository := authRepo.NewAuthRepository(initDB)
	authService := authService.NewAuthService(authRepository, userService, hashing, jwtInterface)
	authHandler := authHandler.NewAuthHandler(authService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to API GOTOPUP",
		})
	})

	// router
	router.BootAuthRoute(app, authHandler)
	addr := fmt.Sprintf(":%d", initConfig.AppPort)
	if err := app.Listen(addr).Error(); err != addr {
		panic("application failed to start")
	}
}
