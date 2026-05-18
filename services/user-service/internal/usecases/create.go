package usecases

import "users/internal/domain"

type CreateUserService struct {
	UserRepository UserRepository
}

func NewCreateUserService(repo UserRepository) *CreateUserService {
	return &CreateUserService{
		UserRepository: repo,
	}
}

func (s *CreateUserService) CreateUser(username, email string) (*domain.User, error) {
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
