package models

import (
	"time"
)

type Pipelines struct {
	Id          string    `xorm:"not null pk comment('ID') CHAR(36)"`
	TeamId      string    `xorm:"not null comment('团队ID') index CHAR(36)"`
	Name        string    `xorm:"not null comment('名称') VARCHAR(255)"`
	Description string    `xorm:"not null comment('描述') VARCHAR(255)"`
	Spec        string    `xorm:"not null comment('定时器') CHAR(64)"`
	Status      int       `xorm:"not null default 1 comment('状态') TINYINT(1)"`
	Finished    string    `xorm:"comment('成功时执行') CHAR(36)"`
	Failed      string    `xorm:"comment('失败时执行') CHAR(36)"`
	Overlap     int       `xorm:"not null default 0 comment('重复执行') TINYINT(1)"`
	CreatedAt   time.Time `xorm:"not null created comment('创建于') TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"not null updated comment('更新于') TIMESTAMP"`
}

// 定义模型的数据表名称
func (pipelines *Pipelines) TableName() string {
	return "pipelines"
}
