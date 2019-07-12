package model

import (
	"time"
)

//
// SecretView type reperesents the db structure of user.
//
type SecretView struct {
	ID        int
	SecretID  int
	CreatedAt time.Time
}
