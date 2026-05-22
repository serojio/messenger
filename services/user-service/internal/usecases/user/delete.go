package user

import (
	"context"
	"users/internal/domain"

	"github.com/google/uuid"
)

type DeleteUseCase struct {
	userRepository UserRepository
}

func NewDeleteUseCase(repo UserRepository) *DeleteUseCase {
	return &DeleteUseCase{
		userRepository: repo,
	}
}

type DeleteCommand struct {
	ID uuid.UUID
}

func (s *DeleteUseCase) Execute(ctx context.Context, cmd DeleteCommand) (*domain.User, error) {
	user, err := s.userRepository.GetByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	err = s.userRepository.DeleteByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// func (s *DeleteUseCase) Delete(ctx context.Context, id uint) (*domain.User, error) {
// 	user, err := s.userRepository.DeleteByID(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }
