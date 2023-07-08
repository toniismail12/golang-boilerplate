package services

import (
	"boilerplate-go/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupCors(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     GetDomain(),
		AllowHeaders:     config.AllowHeaders(),
		AllowMethods:     config.AllowMethods(),
		AllowCredentials: true,
	}))
}
