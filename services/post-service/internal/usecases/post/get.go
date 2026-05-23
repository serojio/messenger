package post

import (
	"context"
	"posts/internal/domain"

	"github.com/google/uuid"
)

type GetUseCase struct {
	postRepository PostRepository
}

func NewGetUseCase(repo PostRepository) *GetUseCase {
	return &GetUseCase{
		postRepository: repo,
	}
}

type GetCommand struct {
	ID uuid.UUID
}

func (s *GetUseCase) Execute(ctx context.Context, cmd GetCommand) (*domain.Post, error) {
	post, err := s.postRepository.GetByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	return post, nil
}
