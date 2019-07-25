package models

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"reflect"
	"time"
)

// 用户操作日志
type Log struct {
	Id        int64     `xorm:"pk autoincr comment('ID') BIGINT(20)"`
	UserId    string    `xorm:"not null comment('用户ID') index CHAR(36)"`
	Operation string    `xorm:"not null comment('操作') VARCHAR(255)"`
	Result    string    `xorm:"null comment('结果') LONGTEXT(0)"`
	CreatedAt time.Time `xorm:"not null comment('创建于') created DATETIME"`
}

// 定义模型的数据表名称
func (log *Log) TableName() string {
	return "logs"
}

func (log *Log) Store() error {
	_, err := Engine.Insert(log)
	return err
}

// Create operation log
func CreateLog(v interface{}, ctx iris.Context, operation string) error {
	token := ctx.Values().Get("jwt").(*jwt.Token)
	claims, _ := token.Claims.(jwt.MapClaims)
	id := claims["sub"]

	var (
		result []byte
		err error
	)

	switch reflect.TypeOf(v).String() {
	case "models.Team":
		obj := reflect.ValueOf(v).Interface().(Team)
		result, err = json.Marshal(&obj)
		if err != nil {
			return err
		}
		break
	case "models.User":
		obj := reflect.ValueOf(v).Interface().(User)
		result, err = json.Marshal(&obj)
		if err != nil {
			return err
		}
		break
	case "models.Task":
		obj := reflect.ValueOf(v).Interface().(User)
		result, err = json.Marshal(&obj)
		if err != nil {
			return err
		}
		break
	case "models.Node":
		obj := reflect.ValueOf(v).Interface().(User)
		result, err = json.Marshal(&obj)
		if err != nil {
			return err
		}
		break
	case "models.Pipeline":
		obj := reflect.ValueOf(v).Interface().(User)
		result, err = json.Marshal(&obj)
		if err != nil {
			return err
		}
		break
	}

	log := &Log{
		UserId: id.(string),
		Operation: operation,
		Result: string(result),
		CreatedAt: time.Now(),
	}

	return log.Store()
}

func (log *Log) MarshalJSON() ([]byte, error) {
	type Alias Log
	return json.Marshal(&struct {
		Alias
		CreatedAt string `json:"created_at"`
	}{
		Alias:     Alias(*log),
		CreatedAt: log.CreatedAt.Format(DefaultTimeFormat),
	})
}