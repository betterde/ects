package models

import "time"

type (
	Team struct {
		Id          string    `xorm:"not null pk comment('ID') CHAR(36)"`
		Name        string    `xorm:"not null comment('名称') VARCHAR(255)"`
		Description string    `xorm:"not null comment('描述') VARCHAR(255)"`
		CreatedAt   time.Time `xorm:"not null created comment('创建于') TIMESTAMP"`
		UpdatedAt   time.Time `xorm:"not null updated comment('更新于') TIMESTAMP"`
		Users       []*User `json:"users" xorm:"-"`
	}
)

// 定义模型的数据表名称
func (team *Team) TableName() string {
	return "teams"
}