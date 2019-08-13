package setting

import (
	"github.com/betterde/ects/internal/response"
	"github.com/betterde/ects/models"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"strconv"
)

type (
	Controller           struct{}
	UpdateSettingRequest struct {
		Url        string `json:"url" validate:"required"`
		Host       string `json:"host" validate:"required"`
		Port       int    `json:"port" validate:"numeric"`
		User       string `json:"user" validate:"required"`
		Pass       string `json:"pass" validate:"required"`
		Name       string `json:"name" validate:"required"`
		Protocol   string `json:"protocol" validate:"required"`
		Encryption string `json:"encryption" validate:"required"`
	}
)

func (instance *Controller) PostMail(ctx iris.Context) mvc.Response {
	return response.Success("", response.Payload{"data": make([]interface{}, 0)})
}

// 更新服务配置
func (instance *Controller) PutNotification(ctx iris.Context) mvc.Response {

	session := models.Engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return response.InternalServerError("开启事务失败", err)
	}

	settings := make([]*models.Configuration, 0)

	if err := session.Find(&settings); err != nil {
		return response.InternalServerError("获取系统设置失败", err)
	}

	// 查看设置是否已存在
	if len(settings) == 0 {
		params := UpdateSettingRequest{}

		if err := ctx.ReadJSON(&params); err != nil {
			return response.Send(400, "解析参数失败", err)
		}

		url := &models.Configuration{
			Name:        "url",
			Value:       params.Url,
			Default:     "",
			Module:      "notification",
			Description: "系统通知中的跳转链接地址",
		}
		settings = append(settings, url)

		host := &models.Configuration{
			Name:        "host",
			Value:       params.Host,
			Default:     "",
			Module:      "notification",
			Description: "邮件服务器主机地址",
		}
		settings = append(settings, host)

		port := &models.Configuration{
			Name:        "port",
			Value:       strconv.Itoa(params.Port),
			Default:     "25",
			Module:      "notification",
			Description: "邮件服务器端口",
		}
		settings = append(settings, port)

		user := &models.Configuration{
			Name:        "user",
			Value:       params.User,
			Default:     "",
			Module:      "notification",
			Description: "邮箱系统账户",
		}
		settings = append(settings, user)

		pass := &models.Configuration{
			Name:        "pass",
			Value:       params.Pass,
			Default:     "",
			Module:      "notification",
			Description: "邮箱系统密码",
		}
		settings = append(settings, pass)

		name := &models.Configuration{
			Name:        "name",
			Value:       params.Name,
			Default:     "",
			Module:      "notification",
			Description: "邮件 From 名称",
		}
		settings = append(settings, name)

		protocol := &models.Configuration{
			Name:        "protocol",
			Value:       params.Protocol,
			Default:     "",
			Module:      "notification",
			Description: "发送邮件协议",
		}
		settings = append(settings, protocol)

		encryption := &models.Configuration{
			Name:        "encryption",
			Value:       params.Encryption,
			Default:     "",
			Module:      "notification",
			Description: "邮件加密类型",
		}
		settings = append(settings, encryption)

		if count, err := session.Insert(settings); err != nil && count != 8 {
			return response.InternalServerError("更新通知设置失败", err)
		}
	} else {
		//params := make(map[string]interface{})
		//if err := ctx.ReadJSON(&params); err != nil {
		//	log.Println(err)
		//	return response.InternalServerError("解析参数失败", err)
		//}
		//for _, setting := range settings {
		//	value := params[setting.Name]
		//	log.Println(value)
		//	switch value.(type) {
		//	case string:
		//		setting.Value = value.(string)
		//		break
		//	case float64:
		//		setting.Value = strconv.FormatInt(value.(int64), 10)
		//		break
		//	}
		//	log.Printf("%#v", setting)
		//	if _, err := session.Where(builder.Eq{"name": setting.Name}).Update(setting); err != nil {
		//		if err := session.Rollback(); err != nil {
		//			log.Println(err)
		//		}
		//		log.Println(err)
		//		return response.InternalServerError("更新失败", err)
		//	}
		//}
	}

	//if err := session.Commit(); err != nil {
	//	return response.InternalServerError("提交事务失败", err)
	//}

	return response.Success("", response.Payload{"data": settings})
}
