package table

import (
	"time"
)

type Users struct {
	Id         uint      `json:"id"`
	Badge      string    `json:"badge"`
	Nama       string    `json:"nama"`
	Password   string    `json:"password"`
	Role_Id    int       `json:"role_id"`
	Email      string    `json:"email"`
	Created_at time.Time `json:"created_at"`
	Created_by string    `json:"created_by"`
	Updated_at time.Time
	Updated_by string
	Deleted_at time.Time
	Deleted_by string
}
