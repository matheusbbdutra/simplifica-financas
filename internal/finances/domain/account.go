package domain

import "time"

type Account struct {
	ID *string
	Name string
	Balance float64
	UserID *string
	CreatedAt string
}
	
func NewAccount(
	name string,
	balance float64,
	userID *string,
) *Account {
	return &Account{
		Name:     name,
		Balance:  balance,
		UserID:   userID,
		CreatedAt: time.Now().Format(time.DateTime),
	}
}

