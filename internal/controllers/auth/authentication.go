package auth

import (
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/services"
	"github.com/betterde/ects/internal/utils/response"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

type (
	AuthenticationController struct {
		Service services.UserService
	}

	SignIn struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	SignInSuccess struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int64  `json:"expires_in"`
	}
)

var validate *validator.Validate

func (instance *AuthenticationController) BeforeActivation(request mvc.BeforeActivation) {
	request.Handle("POST", "/signin", "SignInHandler")
	request.Handle("POST", "/signout", "SignOutHandler")
}

// 用户登录逻辑
func (instance *AuthenticationController) SignInHandler(ctx iris.Context) {
	var params SignIn
	validate = validator.New()
	if err := ctx.ReadJSON(&params); err != nil {
		// TODO Add logger
	}

	err := validate.Struct(params)

	if err != nil {
		if _, err := ctx.JSON(response.ValidationError("用户名和密码不能为空")); err != nil {
			// TODO Add logger
		}
		return
	}

	deadline := time.Now().Add(config.Conf.Auth.TTL).Unix()

	token := instance.Service.Attempt(params.Username, params.Password)

	if _, err := ctx.JSON(response.Success("登录成功", SignInSuccess{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   deadline,
	})); err != nil {
		// TODO Add logger
	}
}

func (instance *AuthenticationController) SignOutHandler(ctx iris.Context) {

}
