package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func GetUID(ctx iris.Context) string {
	token := ctx.Values().Get("jwt").(*jwt.Token)
	claims, _ := token.Claims.(jwt.MapClaims)
	return claims["sub"].(string)
}