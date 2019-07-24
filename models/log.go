package models

import (
	"encoding/json"
	"time"
)

// 用户操作日志
type Log struct {
	Id        int64     `xorm:"pk autoincr comment('ID') BIGINT(20)"`
	UserId    string    `xorm:"not null comment('用户ID') index CHAR(36)"`
	Operation string    `xorm:"not null comment('操作') VARCHAR(255)"`
	Result    string    `xorm:"comment('结果') VARCHAR(255)"`
	CreatedAt time.Time `xorm:"not null comment('创建于') created DATETIME"`
}

// 定义模型的数据表名称
func (log *Log) TableName() string {
	return "logs"
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