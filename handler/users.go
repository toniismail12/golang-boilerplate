package handler

import (
	"boilerplate-go/database"
	constants "boilerplate-go/global-variable"
	"boilerplate-go/response"
	"boilerplate-go/services"
	"boilerplate-go/table"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	DB := database.DB

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	badge := c.Query("badge")
	offset := (page - 1) * limit

	var (
		total_rows int64
		data       []response.GetUsers
	)

	query := DB.Table(constants.TABLE_USERS).
		Select("users.*", "roles.role").
		Joins("JOIN roles ON users.role_id = roles.id").
		Where("badge LIKE ?", "%"+badge+"%").
		Order("id desc")

	query.Count(&total_rows)
	query.Offset(offset).Limit(limit)
	query.Find(&data)

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Data Users",
		"data":    data,
		"meta": fiber.Map{
			"page":  page,
			"limit": limit,
			"total": total_rows,
		},
	})

}

func CreateUsers(c *fiber.Ctx) error {

	DB := database.DB

	var (
		request = new(response.FormUsers)
		users   table.Users
	)

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	users.Badge = request.Badge
	users.Nama = request.Nama
	users.Email = request.Nama
	users.Password, _ = services.HashPassword(request.Password)
	users.Role_Id = request.Role_Id
	users.Created_at = services.TimeNow()
	users.Created_by = services.GetUserLogin(c)

	if err := DB.Create(&users).Error; err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Succes Create Data",
		"request": request,
	})

}
