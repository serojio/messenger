package user

type DeleteUseCase struct {
	userRepository UserRepository
}

func NewDeleteUseCase(repo UserRepository) *DeleteUseCase {
	return &DeleteUseCase{
		userRepository: repo,
	}
}

// func (s *DeleteUseCase) Delete(ctx context.Context, id uint) (*domain.User, error) {
// 	user, err := s.userRepository.DeleteByID(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }
