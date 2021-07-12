package entity

import (
	"time"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	Age       uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}
