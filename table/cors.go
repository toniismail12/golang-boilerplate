package table

import (
	"time"
)

type CorsDomain struct {
	Id         uint      `json:"id"`
	Domain     string    `json:"domain"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time
	Created_by string
	Updated_by string
	Deleted_at time.Time
	Deleted_by string
}
