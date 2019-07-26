package models

import "github.com/betterde/ects/internal/utils"

type TaskRecords struct {
	Id               int64      `xorm:"pk autoincr comment('ID') BIGINT(20)"`
	PipelineRecordId string     `xorm:"not null comment('流水线记录ID') index CHAR(36)"`
	TaskId           string     `xorm:"not null comment('任务ID') index CHAR(36)"`
	NodeId           string     `xorm:"not null comment('节点ID') index CHAR(36)"`
	TaskName         string     `xorm:"not null comment('任务名称') VARCHAR(255)"`
	WorkerName       string     `xorm:"not null comment('节点名称') VARCHAR(255)"`
	Content          string     `xorm:"not null comment('执行内容') TEXT"`
	Mode             string     `xorm:"not null comment('执行方式') VARCHAR(255)"`
	Timeout          int        `xorm:"not null default 0 comment('超时时间') INT(10)"`
	Retries          int        `xorm:"not null default 0 comment('重试次数') TINYINT(3)"`
	Status           string     `xorm:"not null default 'finished' comment('状态') VARCHAR(255)"`
	Result           string     `xorm:"not null comment('执行结果') TEXT"`
	CreatedAt        utils.Time `xorm:"not null created comment('创建于') DATETIME"`
	UpdatedAt        utils.Time `xorm:"not null updated comment('更新于') DATETIME"`
}

// 定义模型的数据表名称
func (records *TaskRecords) TableName() string {
	return "task_records"
}
