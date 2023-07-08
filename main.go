package main

import (
	"boilerplate-go/config"
	"boilerplate-go/database"
	"boilerplate-go/routers"
	"boilerplate-go/services"
)

func main() {

	// database connect
	database.Connect()

	// start fiber
	app := services.CreateApp()

	// router
	routers.Setup(app)

	// get port
	port := config.Env("PORT")

	// run
	app.Listen(":" + port)

}
