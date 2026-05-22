package user

import (
	"context"
	"users/internal/domain"

	"github.com/google/uuid"
)

type UpdateUseCase struct {
	userRepository UserRepository
}

func NewUpdateUseCase(repo UserRepository) *UpdateUseCase {
	return &UpdateUseCase{
		userRepository: repo,
	}
}

type UpdateCommand struct {
	ID       uuid.UUID
	Email    *string
	Username *string
}

func (s *UpdateUseCase) Execute(ctx context.Context, cmd UpdateCommand) (*domain.User, error) {
	user, err := s.userRepository.GetByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domain.ErrUserNotFound
	}
	if cmd.Email != nil {
		user.Email = *cmd.Email
	}
	if cmd.Username != nil {
		user.Username = *cmd.Username
	}

	err = s.userRepository.UpdateByID(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
