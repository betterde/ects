package services

import (
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/models"
	"github.com/betterde/ects/internal/repositories"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
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
	log.Println(username, passwod)
	if username == "" || passwod == "" {
		return ""
	}

	user, err := s.repo.RetrieveByCredentials(username, passwod)

	if err != nil {

	}

	token, err := IssueToken(user)

	if err != nil {

	}

	return token
}

// 为用户颁发Access Token
func IssueToken(user *models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "ects",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(config.Conf.Auth.TTL) * time.Second).Unix(),
		"nbf": time.Now().Unix(),
		"sub": user.ID,
	})

	return token.SignedString([]byte(config.Conf.Auth.Secret))
}
