package model

/*
User Domain Model
*/
type User struct {
	ID        int
	FirstName string
	LastName  string
}

/*
Users Domain Model
*/
type Users []User
