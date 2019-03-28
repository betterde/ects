package account

import (
	"github.com/betterde/ects/internal/services"
	"github.com/betterde/ects/internal/utils/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"log"
)

type (
	ProfileController struct {
		Service services.UserService
	}
	Profile struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Email   string `json:"email"`
		Avatar  string `json:"avatar"`
		GroupId int64  `json:"group_id"`
		Manager bool   `json:"manager"`
	}
)

// 获取用户信息
func (instance *ProfileController) Get(ctx iris.Context) {
	token := ctx.Values().Get("jwt").(*jwt.Token)
	claims, _ := token.Claims.(jwt.MapClaims)
	id := claims["sub"]
	user, err := instance.Service.FindByID(id.(string))
	if err != nil {
		if _, err := ctx.JSON(response.NotFound(err.Error())); err != nil {
			log.Println(err)
		}
		return
	}
	if _, err := ctx.JSON(response.Success("请求成功", response.Payload{"data": &Profile{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Avatar:  user.Avatar,
		GroupId: user.GroupId,
		Manager: user.Manager,
	}})); err != nil {
		log.Println(err)
	}
}

func (instance *ProfileController) Post(ctx iris.Context) {

}
