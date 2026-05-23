package post

import (
	"context"
	"posts/internal/domain"

	"github.com/google/uuid"
)

type PostRepository interface {
	Create(ctx context.Context, post *domain.Post) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Post, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
	UpdateByID(ctx context.Context, post *domain.Post) error

	List(ctx context.Context) ([]*domain.Post, error)
}
