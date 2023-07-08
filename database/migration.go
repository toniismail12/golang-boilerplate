package database

import (
	"boilerplate-go/table"
)

func Migration() {

	DB.AutoMigrate(

		&table.Users{},
		&table.Role{},
		&table.User_login{},
		&table.CorsDomain{},
	)

}
