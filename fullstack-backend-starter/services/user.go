package services

import (
	"github.com/customerio/fullstack-backend-homework/models"
	"github.com/customerio/fullstack-backend-homework/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepository.GetByEmail(email)
}
