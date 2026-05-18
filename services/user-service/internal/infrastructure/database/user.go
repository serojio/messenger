package database

import (
	"users/internal/domain"
	"users/internal/usecases"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(database *gorm.DB) usecases.UserRepository {
	return &GormUserRepository{db: database}
}

func (r *GormUserRepository) Create(user *domain.User) (*domain.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *GormUserRepository) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
