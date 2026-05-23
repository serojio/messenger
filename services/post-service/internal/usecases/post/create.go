package post

import (
	"context"
	"posts/internal/domain"
	"time"

	"github.com/google/uuid"
)

type CreateUseCase struct {
	postRepository PostRepository
}

func NewCreateUseCase(repo PostRepository) *CreateUseCase {
	return &CreateUseCase{
		postRepository: repo,
	}
}

type CreateCommand struct {
	Title    string
	Content  string
	AuthorID uuid.UUID
}

func (s *CreateUseCase) Execute(ctx context.Context, cmd CreateCommand) (*domain.Post, error) {
	post := &domain.Post{
		ID:        uuid.New(),
		Title:     cmd.Title,
		Content:   cmd.Content,
		AuthorID:  cmd.AuthorID,
		CreatedAt: time.Now(),
	}

	err := s.postRepository.Create(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}
