package model

import (
	"time"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	BirthDay  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
