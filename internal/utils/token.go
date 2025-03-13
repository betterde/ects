package utils

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func GetUID(ctx iris.Context) string {
	claims := jwt.Get(ctx).(jwt.Claims)
	return claims.Subject
}
