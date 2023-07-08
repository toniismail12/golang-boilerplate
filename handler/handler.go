package handler

import (
	constants "boilerplate-go/global-variable"

	"github.com/gofiber/fiber/v2"
)

func AppName(c *fiber.Ctx) error {

	c.Status(200)
	return c.JSON(fiber.Map{
		"app_name": constants.APP_NAME,
		"desc":     constants.DESC,
	})

}
