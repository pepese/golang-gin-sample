package model

import "time"

/*
User Domain Model
*/
type User struct {
	ID        int
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

/*
Users Domain Model
*/
type Users []User
