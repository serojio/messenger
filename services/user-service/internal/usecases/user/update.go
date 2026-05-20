package user

type UpdateUseCase struct {
	userRepository UserRepository
}

func NewUpdateUseCase(repo UserRepository) *UpdateUseCase {
	return &UpdateUseCase{
		userRepository: repo,
	}
}

// func (s *UpdateUseCase) Update(ctx context.Context, id uint) (*domain.User, error) {
// 	user, err := s.userRepository.Update(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }
