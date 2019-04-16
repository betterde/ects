package models

import (
	"time"
)

type Pipelines struct {
	Id          string    `json:"id" xorm:"not null pk comment('ID') CHAR(36)"`
	TeamId      string    `json:"team_id" xorm:"not null comment('团队ID') index CHAR(36)"`
	Name        string    `json:"name" xorm:"not null comment('名称') VARCHAR(255)"`
	Description string    `json:"description" xorm:"not null comment('描述') VARCHAR(255)"`
	Spec        string    `json:"spec" xorm:"not null comment('定时器') CHAR(64)"`
	Status      int       `json:"status" xorm:"not null default 1 comment('状态') TINYINT(1)"`
	Finished    string    `json:"finished" xorm:"comment('成功时执行') CHAR(36)"`
	Failed      string    `json:"failed" xorm:"comment('失败时执行') CHAR(36)"`
	Overlap     int       `json:"overlap" xorm:"not null default 0 comment('重复执行') TINYINT(1)"`
	CreatedAt   time.Time `json:"created_at" xorm:"not null created comment('创建于') DATETIME"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"not null updated comment('更新于') DATETIME"`
}

// 定义模型的数据表名称
func (pipelines *Pipelines) TableName() string {
	return "pipelines"
}
