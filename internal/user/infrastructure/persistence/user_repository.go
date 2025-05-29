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

func (r *UserRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	var userModel UserModel
	if err := r.db.Where("email = ?", email).First(&userModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err 
	}

	user := NewUserDomainFromModel(&userModel)
	return user, nil
}

func (r *UserRepositoryImpl) Update(user *domain.User) error {
	userModel := NewUserModelFromDomain(user)
	if err := r.db.Save(userModel).Error; err != nil {
		return err
	}

	return nil
}