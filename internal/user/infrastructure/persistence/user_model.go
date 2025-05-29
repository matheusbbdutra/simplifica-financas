package persistence

import (
	"simplificafinancas/internal/user/domain"
	"time"
)

type UserModel struct {
	ID        uint `json:"id" gorm:"primaryKey;autoIncrement"`
	Name	  string `json:"name" gorm:"not null"`
	Email     string `json:"email" gorm:"not null;unique"`
	Password  string `json:"password" gorm:"not null"`
  	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (UserModel) TableName() string {
	return "users"
}

func NewUserModelFromDomain(user *domain.User) *UserModel {
	var id uint
    if user.ID != nil {
        id = *user.ID
    }

    return &UserModel{
        ID:       id,
        Name:      user.Name,
        Email:     user.Email,
        Password:  user.Password,
        CreatedAt: user.CreatedAt,
    }
}


func NewUserDomainFromModel(user *UserModel) *domain.User {
	return &domain.User{
		ID:        &user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}