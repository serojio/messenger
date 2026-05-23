package post

import (
	"context"
	"posts/internal/domain"
)

type ListUseCase struct {
	postRepository PostRepository
}

func NewListUseCase(repo PostRepository) *ListUseCase {
	return &ListUseCase{
		postRepository: repo,
	}
}

func (s *ListUseCase) Execute(ctx context.Context) ([]*domain.Post, error) {
	posts, err := s.postRepository.List(ctx)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
