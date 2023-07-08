package middleware

import (
	"boilerplate-go/database"
	"boilerplate-go/table"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticate(c *fiber.Ctx) error {

	// Get token auth
	token := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	var user table.User_login
	database.DB.Where("jwt = ?", token).First(&user)
	// log.Println(user.Jwt)

	// claim token, aktif atau kadaluarsa
	if _, err := Parsejwt(user.Jwt); err != nil {

		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated || Access Denied",
		})
	}

	// cek ke database login
	if user.Jwt == "" || user.Jwt != token {

		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated || Access Denied",
		})
	}

	return c.Next()
}
