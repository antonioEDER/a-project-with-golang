package example

import (
	gotoolboxtime "github.com/eucatur/go-toolbox/time"
)

type User struct {
	Id int64 `json:"id" db:"sale_id"`

	CreatedAt gotoolboxtime.Timestamp `json:"created_at" db:"sale_created_at"`
}
