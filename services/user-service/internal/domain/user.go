package domain

import "time"

type User struct {
	ID        uint
	Username  string
	Email     string
	CreatedAt time.Time
}
