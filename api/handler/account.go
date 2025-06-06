package handler

import (
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Profile struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	TeamId  string `json:"team_id"`
	Manager bool   `json:"manager"`
}

func GetProfile(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	userService := service.NewUserService()
	id, err := claims.GetSubject()
	if err != nil {
		return err
	}

	user, err := userService.FindByID(id)
	if err != nil {
		return err
	}
	return ctx.JSON(response.Success("请求成功", Profile{
		ID:      user.Id,
		Name:    user.Name,
		Email:   user.Email,
		Manager: user.Manager,
	}, nil))
}
