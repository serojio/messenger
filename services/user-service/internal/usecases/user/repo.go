package user

import (
	"context"
	"users/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	// GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	// DeleteByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	// Update(ctx context.Context, id uuid.UUID) (*domain.User, error)
}
