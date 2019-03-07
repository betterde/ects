package services

import (
	"github.com/betterde/ects/internal/models"
	"github.com/betterde/ects/internal/repositories"
)

type UserService interface {
	Users() []models.User
	Attempt(username, password string) string
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

// 验证用户凭证
func (s *userService) Attempt(username, passwod string) string {
	if username == "" || passwod == "" {
		return ""
	}

	user, err := s.repo.RetrieveByCredentials(username, passwod)

	if err != nil {

	}

	token, err := user.IssueToken()

	if err != nil {

	}

	return token
}