package models

import (
	"time"
)

type PipelineTaskPivot struct {
	PipelineId string    `xorm:"not null comment('ID') index CHAR(36)"`
	TaskId     string    `xorm:"not null comment('ID') index CHAR(36)"`
	Step       int       `xorm:"not null comment('步骤') SMALLINT(5)"`
	Timeout    int       `xorm:"not null default 0 comment('超时时间') INT(10)"`
	Interval   int       `xorm:"not null default 0 comment('间隔时间') INT(10)"`
	Retries    int       `xorm:"not null default 0 comment('重试次数') TINYINT(3)"`
	Dependence string    `xorm:"not null default 'strong' comment('依赖') VARCHAR(255)"`
	CreatedAt  time.Time `xorm:"not null created comment('创建于') DATETIME"`
	UpdatedAt  time.Time `xorm:"not null updated comment('更新于') DATETIME"`
}

// 定义模型的数据表名称
func (pivot *PipelineTaskPivot) TableName() string {
	return "pipeline_task_pivot"
}
