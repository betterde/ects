package models

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"reflect"
	"time"
)

type Log struct {
	Id        int64     `json:"id" xorm:"pk autoincr comment('ID') BIGINT(20)"`
	UserId    string    `json:"user_id" xorm:"not null comment('用户ID') index CHAR(36)"`
	Operation string    `json:"operation" xorm:"not null comment('操作') VARCHAR(255)"`
	Result    string    `json:"result" xorm:"null comment('结果') LONGTEXT(0)"`
	CreatedAt time.Time `json:"created_at" xorm:"not null comment('创建于') created DATETIME"`
}

// Define table name
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
		break
	case "models.User":
		obj := reflect.ValueOf(v).Interface().(User)
		result, err = json.Marshal(&obj)
		break
	case "models.Task":
		obj := reflect.ValueOf(v).Interface().(Task)
		result, err = json.Marshal(&obj)
		break
	case "models.Node":
		obj := reflect.ValueOf(v).Interface().(Node)
		result, err = json.Marshal(&obj)
		break
	case "models.Pipeline":
		obj := reflect.ValueOf(v).Interface().(Pipeline)
		result, err = json.Marshal(&obj)
		break
	}

	if err != nil {
		return err
	}

	log := &Log{
		UserId:    id.(string),
		Operation: operation,
		Result:    string(result),
		CreatedAt: time.Now(),
	}

	return log.Store()
}

// Marshal struct to json
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
