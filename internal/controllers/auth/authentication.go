package auth

import (
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/models"
	"github.com/betterde/ects/internal/services"
	"github.com/betterde/ects/internal/utils/response"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"gopkg.in/go-playground/validator.v9"
	"log"
)

type (
	AuthenticationController struct {
		Service services.UserService
	}

	SignIn struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	SignUp struct {
		Name     string `json:"name" validate:"required"`
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		Confirm  string `json:"confirm" validate:"required"`
	}

	SignInSuccess struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int64  `json:"expires_in"`
	}
)


func (instance *AuthenticationController) BeforeActivation(request mvc.BeforeActivation) {
	request.Handle("POST", "/signin", "SignInHandler")
	request.Handle("POST", "/signout", "SignOutHandler")
}

// 用户登录逻辑
func (instance *AuthenticationController) SignInHandler(ctx iris.Context) {
	var params SignIn
	validate := validator.New()
	if err := ctx.ReadJSON(&params); err != nil {
		// TODO Add logger
	}

	err := validate.Struct(params)

	if err != nil {
		ctx.StatusCode(iris.StatusUnprocessableEntity)
		if _, err := ctx.JSON(response.ValidationError("用户名和密码不能为空")); err != nil {
			log.Println(err)
		}
		return
	}

	token, err := instance.Service.Attempt(params.Username, params.Password)
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		if _, err := ctx.JSON(response.UnAuthenticated(err.Error())); err != nil {
			log.Println(err)
		}
		return
	}

	if _, err := ctx.JSON(response.Success("登录成功", SignInSuccess{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   config.Conf.Auth.TTL,
	})); err != nil {
		// TODO Add logger
	}
}

// 用户注销登陆
func (instance *AuthenticationController) SignOutHandler(ctx iris.Context) {

}

// 用户注册
func (instance *AuthenticationController) SignUpHandler(ctx iris.Context) {
	var params SignUp
	validate := validator.New()
	if err := ctx.ReadJSON(&params); err != nil {
		// TODO Add logger
	}

	if err := validate.Struct(params); err != nil {
		if _, err := ctx.JSON(response.ValidationError("用户名和密码不能为空")); err != nil {
			// TODO Add logger
		}
		return
	}

	user := &models.User{
		Name: params.Name,
	}

	hash, err := models.GeneratePassword(params.Password)
	if err != nil {
		// TODO
	}
	user.Password = string(hash)
	user.Save()
}
