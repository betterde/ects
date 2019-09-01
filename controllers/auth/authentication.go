package auth

import (
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"gopkg.in/go-playground/validator.v9"
	"log"
)

type (
	Controller struct {
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

// 路由分发
func (instance *Controller) BeforeActivation(request mvc.BeforeActivation) {
	request.Handle("POST", "/signin", "SignInHandler")
	request.Handle("POST", "/signout", "SignOutHandler")
}

// 用户登录逻辑
func (instance *Controller) SignInHandler(ctx iris.Context) mvc.Response {
	var params SignIn
	validate := validator.New()
	if err := ctx.ReadJSON(&params); err != nil {
		// TODO Add logger
	}

	err := validate.Struct(params)

	if err != nil {
		return response.ValidationError("用户名和密码不能为空")
	}

	token, err := instance.Service.Attempt(params.Username, params.Password)
	if err != nil {
		return response.UnAuthenticated(err.Error())
	}

	return response.Success("登录成功", response.Payload{"data": SignInSuccess{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   config.Conf.Auth.TTL,
	}})
}

// 用户注销逻辑
func (instance *Controller) SignOutHandler(ctx iris.Context) mvc.Response {
	return response.Success("", response.Payload{"data": make([]interface{}, 0)})
}

// 用户注册逻辑
func (instance *Controller) SignUpHandler(ctx iris.Context) mvc.Response {
	var params SignUp
	validate := validator.New()
	if err := ctx.ReadJSON(&params); err != nil {
		// TODO Add logger
	}

	if err := validate.Struct(params); err != nil {
		return response.ValidationError("用户名和密码不能为空")
	}

	user := &models.User{
		Name: params.Name,
	}

	hash, err := models.GeneratePassword(params.Password)
	if err != nil {
		// TODO
	}
	user.Password = string(hash)

	if err := user.Save(); err != nil {
		log.Println(err)
	}
	return response.Success("注册成功", response.Payload{"data": user})
}
