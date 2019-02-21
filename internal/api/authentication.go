package api

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
)

type (
	AuthenticationController struct {

	}

	SignIn struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
)

func (instance * AuthenticationController) BeforeActivation(request mvc.BeforeActivation)  {
	request.Handle("POST", "/signin", "SignInHandler")
}

// 用户登录逻辑
func (instance * AuthenticationController) SignInHandler(ctx iris.Context) {
	var params SignIn
	if err := ctx.ReadJSON(&params); err != nil {
		
	}
	log.Println(params)
	if _, err := ctx.JSON(ctx.FormValues()); err != nil {
		log.Println(err)
	}
}
