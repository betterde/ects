package utils

import (
	"github.com/kataras/iris"
)

func Pagination(ctx iris.Context) (page, limit, start int) {
	page = ctx.Params().GetIntDefault("page", 1)
	limit = ctx.Params().GetIntDefault("limit", 10)
	start = (page - 1) * limit

	return
}
