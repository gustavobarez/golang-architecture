package domain

import "time"

type User struct {
	Id         int
	FirstName  string
	LastName   string
	Age        int
	CreatedAt  time.Time
	UpdateddAt time.Time
	DeletedAt  time.Time
}
