package user

type GetUseCase struct {
	userRepository UserRepository
}

func NewGetUseCase(repo UserRepository) *GetUseCase {
	return &GetUseCase{
		userRepository: repo,
	}
}

// func (s *GetUseCase) GetByID(ctx context.Context, id uint) (*domain.User, error) {
// 	user, err := s.userRepository.GetByID(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }
