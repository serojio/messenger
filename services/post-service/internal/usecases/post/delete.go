package post

import (
	"context"
	"posts/internal/domain"

	"github.com/google/uuid"
)

type DeleteUseCase struct {
	postRepository PostRepository
}

func NewDeleteUseCase(repo PostRepository) *DeleteUseCase {
	return &DeleteUseCase{
		postRepository: repo,
	}
}

type DeleteCommand struct {
	ID uuid.UUID
}

func (s *DeleteUseCase) Execute(ctx context.Context, cmd DeleteCommand) (*domain.Post, error) {
	post, err := s.postRepository.GetByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, domain.ErrPostNotFound
	}

	err = s.postRepository.DeleteByID(ctx, post.ID)
	if err != nil {
		return nil, err
	}

	return post, nil
}
