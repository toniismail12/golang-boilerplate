package response

import (
	"time"
)

type GetUsers struct {
	Id         uint      `json:"id"`
	Badge      string    `json:"badge"`
	Nama       string    `json:"nama"`
	Role_Id    int       `json:"role_id"`
	Role       string    `json:"role"`
	Email      string    `json:"email"`
	Created_at time.Time `json:"created_at"`
	Created_by string    `json:"created_by"`
}

type FormUsers struct {
	Badge    string `json:"badge"`
	Nama     string `json:"nama"`
	Role_Id  int    `json:"role_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
