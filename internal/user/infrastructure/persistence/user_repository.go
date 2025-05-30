package persistence

import (
	"simplificafinancas/internal/user/domain"

	"gorm.io/gorm"
)


type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		Db: db,
	}
}

func (r *UserRepositoryImpl) Create(user *domain.User) error {
	userModel := NewUserModelFromDomain(user)
	if err := r.Db.Create(userModel).Error; err != nil {
		return err
	}

    return nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	var userModel UserModel
	if err := r.Db.Where("email = ?", email).First(&userModel).Error; err != nil {
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
	if err := r.Db.Save(userModel).Error; err != nil {
		return err
	}

	return nil
}