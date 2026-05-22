package user

import (
	"context"
	"users/internal/domain"

	"github.com/google/uuid"
)

type GetUseCase struct {
	userRepository UserRepository
}

func NewGetUseCase(repo UserRepository) *GetUseCase {
	return &GetUseCase{
		userRepository: repo,
	}
}

type GetCommand struct {
	ID uuid.UUID
}

func (s *GetUseCase) Execute(ctx context.Context, cmd GetCommand) (*domain.User, error) {
	user, err := s.userRepository.GetByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
