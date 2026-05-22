package user

import (
	"context"
	"users/internal/domain"
)

type ListUseCase struct {
	userRepository UserRepository
}

func NewListUseCase(repo UserRepository) *ListUseCase {
	return &ListUseCase{
		userRepository: repo,
	}
}

func (s *ListUseCase) Execute(ctx context.Context) ([]*domain.User, error) {
	users, err := s.userRepository.List(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
