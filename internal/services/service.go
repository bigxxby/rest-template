package services

import (
	"app/internal/models"
	"app/internal/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{UserRepository: repo}
}
func (s *UserService) GetUserById(id int) (*models.User, error) {
	//some buisness logic
	s.UserRepository.GetUserByID(id)
	return nil, nil
}
