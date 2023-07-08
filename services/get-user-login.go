package services

import (
	"boilerplate-go/database"
	"boilerplate-go/table"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetUserLogin(c *fiber.Ctx) string {

	token := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	var user table.User_login
	database.DB.Where("jwt = ?", token).First(&user)

	return user.Username

}
