package middleware

import "github.com/kataras/iris"

func AuthGuard(ctx iris.Context)  {
	// 用户登录认证中间件逻辑
	ctx.Next()
}