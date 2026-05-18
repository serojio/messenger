package usecases

import "users/internal/domain"

type UserService struct {
	UserRepository UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		UserRepository: repo,
	}
}

func (s *UserService) CreateUser(username, email string) (*domain.User, error) {
	user := &domain.User{
		Username: username,
		Email:    email,
	}

	user, err := s.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
