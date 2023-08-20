package user

import (
	"time"
)

type User struct {
	ID        uint
	Username  string
	Email     string
	About     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
