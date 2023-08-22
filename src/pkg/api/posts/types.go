package posts

import (
	"time"

	"github.com/Cknappen/Knappenblog-be/src/pkg/api/user"
)

type Post struct {
	ID        uint
	Title     string
	Publisher user.User
	Tags      string
	Links     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
