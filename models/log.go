package models

import (
	"encoding/json"
	"time"
)

type Log struct {
	Id        int64     `json:"id" xorm:"pk autoincr comment('ID') BIGINT(20)"`
	UserId    string    `json:"user_id" xorm:"not null comment('用户ID') index CHAR(36)"`
	Operation string    `json:"operation" xorm:"not null comment('操作') VARCHAR(255)"`
	Result    string    `json:"result" xorm:"null comment('结果') LONGTEXT(0)"`
	CreatedAt time.Time `json:"created_at" xorm:"not null comment('创建于') created DATETIME"`
}

// 定义日志表名称
func (log *Log) TableName() string {
	return "logs"
}

// 保存日志
func (log *Log) Store() error {
	_, err := Engine.Insert(log)
	return err
}

// 创建用户操作日志
func CreateLog(model Model, uid string, operation string) error {
	var (
		result string
		err error
	)

	result, err = model.ToString()
	if err != nil {
		return err
	}

	log := &Log{
		UserId:    uid,
		Operation: operation,
		Result:    result,
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
