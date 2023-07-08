package routers

import (
	"boilerplate-go/handler"
	"boilerplate-go/middleware"
	"boilerplate-go/services"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// setup cors
	services.SetupCors(app)

	app.Get("/", handler.AppName)

	app.Post("/login", handler.Auth)

	// route dibawah ini wajib login
	app.Use(middleware.IsAuthenticate)

	app.Get("/logout", handler.Logout)

	// users
	app.Post("/users", handler.CreateUsers)
	app.Get("/users", handler.GetUsers)

}
