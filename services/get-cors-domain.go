package services

import (
	"boilerplate-go/database"
	constants "boilerplate-go/global-variable"
	"boilerplate-go/table"
	"log"
)

func GetDomain() string {

	db := database.DB

	var (
		data   []table.CorsDomain
		result string
	)

	if err := db.Table(constants.TABLE_CORS_DOMAIN).Find(&data).Error; err != nil {
		log.Println("Failed to retrieve data from the database:", err)
	}

	for i, domain := range data {
		if i > 0 {
			result += ", "
		}
		result += domain.Domain
	}

	return result

}
