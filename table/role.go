package table

import "time"

type Role struct {
	Id         uint `json:"id"`
	Role       string
	Created_at time.Time
	Created_by string
	Deleted_at time.Time
	Deleted_by string
}
