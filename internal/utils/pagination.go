package utils

import (
	"github.com/kataras/iris"
)

func Pagination(ctx iris.Context) (page, limit, start int) {
	page = ctx.URLParamIntDefault("page", 1)
	limit = ctx.URLParamIntDefault("limit", 10)
	start = (page - 1) * limit

	return
}
