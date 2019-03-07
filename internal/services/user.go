package services

import (
	"github.com/betterde/ects/internal/models"
	"github.com/betterde/ects/internal/repositories"
)

type UserService interface {
	Users() []models.User
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

type userService struct {
	repo repositories.UserRepository
}

func (s *userService) Users() []models.User {
	return s.repo.SelectMany(func(_ models.User) bool {
		return true
	}, -1)
}