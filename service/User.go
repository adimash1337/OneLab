package service

import (
	"awesomeProject/model"
	"awesomeProject/storage"
)

type UserService struct {
	userRepository storage.UserRepository
}

func NewUserService(userRepository storage.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.userRepository.Create(user)
}

func (s *UserService) GetUserByID(id int) (*model.User, error) {
	return s.userRepository.GetByID(id)
}
