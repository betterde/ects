package auth

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
	request.Handle("POST", "/signout", "SignOutHandler")
}

// 用户登录逻辑
func (instance * AuthenticationController) SignInHandler(ctx iris.Context) {
	var params SignIn
	if err := ctx.ReadJSON(&params); err != nil {
		
	}

	if params.Username == "" || params.Password == "" {
		//return ctx.JSON()
	}

	log.Println(params)
	if _, err := ctx.JSON(params); err != nil {
		log.Println(err)
	}
}

func (instance * AuthenticationController) SignOutHandler(ctx iris.Context) {

}
