package services

import (
	"errors"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-xorm/builder"
	"log"
	"strconv"
	"time"
)

type (
	UserInterface interface {
		Users(params map[string]string) (*[]models.User, *response.Meta)
		FindByID(id string) (*models.User, error)
		FindByEmail(email string) (*models.User, error)
		Attempt(username, password string) (string, error)
		Destroy(id string, force bool) error
	}

	UserService struct {
	}
)

func NewUserService() UserInterface {
	return &UserService{}
}

// 获取用户信息
func (service *UserService) Users(params map[string]string) (*[]models.User, *response.Meta) {
	var (
		page  = 1
		limit = 10
		start int
		total int64
		err   error
	)

	if value, exist := params["page"]; exist == true {
		v, err := strconv.Atoi(value)
		if err != nil {

		}
		if v >= 0 {
			page = v
		}
	}

	if value, exist := params["limit"]; exist == true {
		v, err := strconv.Atoi(value)
		if err != nil {

		}
		if v >= 0 {
			limit = v
		}
	}

	start = (page - 1) * limit
	users := make([]models.User, 0)

	if search, exist := params["search"]; exist && search != "" {
		total, err = models.Engine.Where(builder.Like{"name", search}).Count(&models.User{})
		err = models.Engine.Where(builder.Like{"name", search}).Limit(limit, start).Find(&users)
	} else {
		total, err = models.Engine.Count(&models.User{})
		err = models.Engine.Limit(limit, start).Find(&users)
	}

	if err != nil {
		log.Println(err)
	}

	return &users, &response.Meta{
		Limit: limit,
		Page:  page,
		Total: int(total),
	}
}

// 验证用户凭证
func (service *UserService) Attempt(username, passwod string) (token string, err error) {
	user, err := service.RetrieveByCredentials(username, passwod)

	if err != nil {
		return "", err
	}

	token, err = IssueToken(user)
	return token, err
}

// 根据用户凭证获取用户模型
func (service *UserService) RetrieveByCredentials(username, password string) (user *models.User, err error) {
	user, err = service.FindByEmail(username)

	if err != nil {
		return nil, err
	}

	result, err := models.ValidatePassword(password, []byte(user.Password))

	if result {
		return user, err
	}

	return nil, errors.New("用户名或密码错误")
}

// 根据用户邮箱查询用户信息
func (service *UserService) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result, err := models.Engine.Unscoped().Where(builder.Eq{"email": email}).Get(&user)

	if result {
		// 如果用户已经被删除则
		if user.DeletedAt.Unix() != 0 {
			return &user, errors.New("用户已禁用")
		}
		return &user, err
	}

	return &user, errors.New("用户不存在")
}

// 根据用户ID获取用户信息
func (service *UserService) FindByID(id string) (*models.User, error) {
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

// 删除用户信息
func (service *UserService) Destroy(id string, force bool) (err error) {
	var result int64
	if force {
		result, err = models.Engine.Id(id).Unscoped().Delete(&models.User{})
	} else {
		result, err = models.Engine.Id(id).Delete(&models.User{})
	}

	if err != nil && result > 0 {
		return nil
	}

	return errors.New("用户删除失败")
}

// 为用户颁发Access Token
func IssueToken(user *models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "ects",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(config.Conf.Auth.TTL) * time.Second).Unix(),
		"nbf": time.Now().Unix(),
		"sub": user.Id,
	})

	return token.SignedString([]byte(config.Conf.Auth.Secret))
}
