package main

import (
	"fmt"
	"github.com/RianIhsan/go-topup-midtrans/config"
	"github.com/RianIhsan/go-topup-midtrans/utils/db"
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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to API GOTOPUP",
		})
	})
	addr := fmt.Sprintf(":%d", initConfig.AppPort)
	if err := app.Listen(addr).Error(); err != addr {
		panic("Appilaction failed to start")
	}
}
