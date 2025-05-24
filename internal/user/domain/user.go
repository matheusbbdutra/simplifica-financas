package domain

import "time"

type User struct {
	ID *string
	Name string
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func NewUser(
	name string,
	email string,
	password string,
) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		CreatedAt: time.Now(),
	}
}