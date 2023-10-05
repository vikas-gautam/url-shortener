package models

import "time"

type DBUser struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	IsActive  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
