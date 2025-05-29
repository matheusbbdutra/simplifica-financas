package domain

import (
	"simplificafinancas/pkg/utils"
	"time"
)

type User struct {
	ID *uint
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

func (u *User) CheckPassword(password string) bool {
	return utils.CheckPasswordHash(password, u.Password)
}