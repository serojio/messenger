package usecases

import "users/internal/domain"

type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
	GetByID(id uint) (*domain.User, error)
}
