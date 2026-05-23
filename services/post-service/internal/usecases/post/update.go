package post

import (
	"context"
	"posts/internal/domain"

	"github.com/google/uuid"
)

type UpdateUseCase struct {
	postRepository PostRepository
}

func NewUpdateUseCase(repo PostRepository) *UpdateUseCase {
	return &UpdateUseCase{
		postRepository: repo,
	}
}

type UpdateCommand struct {
	ID      uuid.UUID
	Title   *string
	Content *string
}

func (s *UpdateUseCase) Execute(ctx context.Context, cmd UpdateCommand) (*domain.Post, error) {
	post, err := s.postRepository.GetByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, domain.ErrPostNotFound
	}
	if cmd.Title != nil {
		post.Title = *cmd.Title
	}
	if cmd.Content != nil {
		post.Content = *cmd.Content
	}

	err = s.postRepository.UpdateByID(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}
