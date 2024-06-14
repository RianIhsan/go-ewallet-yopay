package main

import (
	"fmt"
	"github.com/RianIhsan/go-topup-midtrans/config"
	authHandler "github.com/RianIhsan/go-topup-midtrans/feature/auth/handler"
	authRepo "github.com/RianIhsan/go-topup-midtrans/feature/auth/repository"
	authService "github.com/RianIhsan/go-topup-midtrans/feature/auth/service"
	topupHandler "github.com/RianIhsan/go-topup-midtrans/feature/topup/handler"
	topupRepo "github.com/RianIhsan/go-topup-midtrans/feature/topup/repository"
	topupService "github.com/RianIhsan/go-topup-midtrans/feature/topup/service"
	userHandler "github.com/RianIhsan/go-topup-midtrans/feature/users/handler"
	userRepo "github.com/RianIhsan/go-topup-midtrans/feature/users/repository"
	userService "github.com/RianIhsan/go-topup-midtrans/feature/users/service"
	"github.com/RianIhsan/go-topup-midtrans/router"
	"github.com/RianIhsan/go-topup-midtrans/utils/db"
	generator "github.com/RianIhsan/go-topup-midtrans/utils/genrator"
	"github.com/RianIhsan/go-topup-midtrans/utils/hashing"
	"github.com/RianIhsan/go-topup-midtrans/utils/jwtToken"
	"github.com/RianIhsan/go-topup-midtrans/utils/payment"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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
	jwtInterface := jwtToken.NewJWT(initConfig.Secret)
	hashing := hashing.NewHash()
	coreApi := payment.InitMidtransCore(*initConfig)
	generatorId := generator.NewGeneratorUUID(initDB)

	cld, err := cloudinary.NewFromParams(initConfig.Cloudinary.CLoudiName, initConfig.Cloudinary.CloudiKey, initConfig.Cloudinary.CloudiSecret)
	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary, %v", err)
	}

	qrCodeGen := generator.NewQrGenerator(cld)

	userRepository := userRepo.NewUserRepository(initDB)
	userService := userService.NewUserService(userRepository)
	userHandler := userHandler.NewUserHandler(userService)

	authRepository := authRepo.NewAuthRepository(initDB)
	authService := authService.NewAuthService(authRepository, userService, hashing, jwtInterface, qrCodeGen)
	authHandler := authHandler.NewAuthHandler(authService)

	topUpRepository := topupRepo.NewTopUpRepository(initDB, coreApi)
	topUpService := topupService.NewTopUpService(topUpRepository, userService, generatorId)
	topUpHandler := topupHandler.NewTopUpHandler(topUpService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to API GOTOPUP",
		})
	})

	// router
	router.BootAuthRoute(app, authHandler)
	router.BootUserRoute(app, userHandler, jwtInterface, userService)
	router.BootTopUpBalanceRoute(app, topUpHandler, jwtInterface, userService)
	addr := fmt.Sprintf(":%d", initConfig.AppPort)
	if err := app.Listen(addr).Error(); err != addr {
		panic("application failed to start")
	}
}
