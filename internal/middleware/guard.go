package middleware

import (
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/response"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"log"
)

var (
	JWTHandler = jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Conf.Auth.Secret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: func(ctx iris.Context, s string) {
			ctx.StatusCode(iris.StatusUnauthorized)
			if _, err := ctx.JSON(response.Response{
				Code:    iris.StatusUnauthorized,
				Message: "Unauthenticated.",
				Data:    make(map[string]interface{}),
			}); err != nil {
				log.Println(err)
			}
		},
	})
)
