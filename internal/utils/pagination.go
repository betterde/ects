package utils

import (
	"github.com/betterde/ects/internal/response"
	"github.com/gofiber/fiber/v2"
)

func Pagination(ctx *fiber.Ctx) (meta response.Meta) {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	meta = response.Meta{
		Page:  ctx.QueryInt("page", 1),
		Limit: ctx.QueryInt("limit", 10),
		Total: 0,
		Start: (page - 1) * limit,
	}
	return
}
