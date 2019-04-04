package account

import (
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/internal/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type (
	Controller struct {
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
func (instance *Controller) Get(ctx iris.Context) mvc.Result {
	token := ctx.Values().Get("jwt").(*jwt.Token)
	claims, _ := token.Claims.(jwt.MapClaims)
	id := claims["sub"]
	user, err := instance.Service.FindByID(id.(string))
	if err != nil {
		return response.NotFound(err.Error())
	}
	return response.Success("请求成功", response.Payload{"data": &Profile{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Avatar:  user.Avatar,
		GroupId: user.GroupId,
		Manager: user.Manager,
	}})
}

func (instance *Controller) Post(ctx iris.Context) {

}
