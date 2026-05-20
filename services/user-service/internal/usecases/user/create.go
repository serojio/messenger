package user

import (
	"context"
	"time"
	"users/internal/domain"

	"github.com/google/uuid"
)

type CreateUseCase struct {
	userRepository UserRepository
}

func NewCreateUseCase(repo UserRepository) *CreateUseCase {
	return &CreateUseCase{
		userRepository: repo,
	}
}

type CreateCommand struct {
	Username string
	Email    string
}

func (s *CreateUseCase) Execute(ctx context.Context, cmd CreateCommand) (*domain.User, error) {
	user := &domain.User{
		ID:        uuid.New(),
		Username:  cmd.Username,
		Email:     cmd.Email,
		CreatedAt: time.Now(),
	}

	err := s.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
