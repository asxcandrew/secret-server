package model

import (
	"time"
)

//
// Secret type reperesents the db structure of user.
//
type Secret struct {
	ID         int
	Body       string
	Hash       string
	ViewsLimit int
	ExpiresAt  time.Time
	CreatedAt  time.Time
}
