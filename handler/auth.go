package handler

import (
	"boilerplate-go/database"
	"boilerplate-go/middleware"
	"boilerplate-go/response"
	"boilerplate-go/services"
	"boilerplate-go/table"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {

	var request = new(response.ReqAuth)

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	// validasi auth pusri
	StatusCode := services.Auth_pusri_validation(request.Username, request.Password)

	// kondisi success validasi
	if StatusCode == 200 {

		// save user login
		services.Save_user_login(request.Username)

		var user table.User_login
		database.DB.Where("username=?", request.Username).First(&user)

		token, err := middleware.GenerateJwt(strconv.Itoa(int(user.Id)))

		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		// mengisi token ke kolom jwt
		isi_jwt := table.User_login{
			Jwt: token,
		}
		database.DB.Where("username=?", request.Username).Updates(&isi_jwt)

		c.Status(200)
		return c.JSON(fiber.Map{
			"message":      "Berhasil Login",
			"username":     request.Username,
			"role":         user.Role,
			"nama":         user.Name,
			"departemen":   user.Departemen,
			"access_token": token,
			"type_token":   "jwt",
			"auth_type":    "Bearer",
		})
	}

	// kondisi gagal validasi
	if StatusCode != 200 {

		var user table.Users
		database.DB.Where("badge=?", request.Username).First(&user)

		// save user login
		services.Save_user_login_local(request.Username)

		password := user.Password // bcrypt password user
		hash := request.Password  // inputan password

		match := services.CheckPasswordHash(password, hash)

		if user.Id == 0 {
			c.Status(401)
			return c.JSON(fiber.Map{
				"message": "username atau password salah",
			})
		}

		if !match {
			c.Status(401)
			return c.JSON(fiber.Map{
				"message": "password atau username salah",
			})
		}

		var user2 table.User_login
		database.DB.Where("username=?", request.Username).First(&user2)

		token, err := middleware.GenerateJwt(strconv.Itoa(int(user2.Id)))

		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		// mengisi token ke kolom jwt
		isi_jwt := table.User_login{
			Jwt: token,
		}
		database.DB.Where("username=?", request.Username).Updates(&isi_jwt)

		c.Status(200)
		return c.JSON(fiber.Map{
			"message":      "Berhasil Login",
			"username":     request.Username,
			"role":         user2.Role,
			"nama":         user.Nama,
			"departemen":   user2.Departemen,
			"access_token": token,
			"type_token":   "jwt",
			"auth_type":    "Bearer",
		})
	}

	c.Status(401)
	return c.JSON(fiber.Map{
		"message": "Gagal Login",
	})

}

func Logout(c *fiber.Ctx) error {

	db := database.DB

	data := table.User_login{
		Jwt:          "-",
		Lates_logout: services.TimeNow(),
	}
	db.Where("username=?", services.GetUserLogin(c)).Updates(&data)

	return c.JSON(fiber.Map{
		"message": "Berhasil Logout",
	})

}
