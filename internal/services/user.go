package services

import (
	"errors"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-xorm/builder"
	"log"
	"time"
)

type UserService interface {
	Users() []models.User
	FindByID(id string) (*models.User, error)
	Attempt(username, password string) (string, error)
}

func NewUserService() UserService {
	return &userService{}
}

type userService struct {
}

func (service *userService) Users() []models.User {
	return []models.User{}
}

// 验证用户凭证
func (service *userService) Attempt(username, passwod string) (token string, err error) {
	user, err := service.RetrieveByCredentials(username, passwod)
	token, err = IssueToken(user)
	return token, err
}

// 根据用户凭证获取用户模型
func (service *userService) RetrieveByCredentials(username, password string) (user *models.User, err error) {
	user, err = service.FindByEmail(username)
	result, err := models.ValidatePassword(password, []byte(user.Password))

	if result {
		return user, err
	}

	return nil, errors.New("用户名或密码错误")
}

// 根据用户邮箱查询用户信息
func (service *userService) FindByEmail(email string) (user *models.User, err error) {
	_, err = models.Engine.Where(builder.Eq{"email": email}).Get(user)
	return
}

// 根据用户ID获取用户信息
func (service *userService) FindByID(id string) (*models.User, error) {
	var user models.User
	result, err := models.Engine.Id(id).Get(&user)
	if err != nil {
		log.Println(err)
	}

	if result {
		return &user, nil
	}

	return &user, errors.New("用户不存在")
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
