package models

import "time"

type User struct {
	Username  string
	Email     string
	Admin     bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
