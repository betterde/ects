package middleware

import (
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/utils/response"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
)

var (
	JWTHandler = jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Conf.Auth.Secret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: func(ctx iris.Context, s string) {
			if _, err := ctx.JSON(response.UnAuthenticated("认证失败请重新登陆")); err != nil {
				// TODO Add loger
			}
		},
	})
)
