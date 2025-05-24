package persistence

import (
	"simplificafinancas/internal/user/domain"

	"gorm.io/gorm"
)


type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) Create(user *domain.User) error {
	userModel := NewUserModelFromDomain(user)
	if err := r.db.Create(userModel).Error; err != nil {
		return err
	}

    return nil
}