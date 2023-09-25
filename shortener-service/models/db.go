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

type DBURL struct {
	ID        int
	UserID    int
	ActualURL string
	ShortURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DbCounter struct {
	ID         int
	UserID     int
	ShortURL   string
	HitCounter int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
