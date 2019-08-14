package models

import "github.com/betterde/ects/internal/utils"

type TaskRecords struct {
	Id               int64      `json:"id" xorm:"pk autoincr comment('ID') BIGINT(20)"`
	PipelineRecordId string     `json:"pipeline_record_id" xorm:"not null comment('流水线记录ID') index CHAR(36)"`
	TaskId           string     `json:"task_id" xorm:"not null comment('任务ID') index CHAR(36)"`
	NodeId           string     `json:"node_id" xorm:"not null comment('节点ID') index CHAR(36)"`
	TaskName         string     `json:"task_name" xorm:"not null comment('任务名称') VARCHAR(255)"`
	WorkerName       string     `json:"worker_name" xorm:"not null comment('节点名称') VARCHAR(255)"`
	Content          string     `json:"content" xorm:"not null comment('执行内容') TEXT"`
	Mode             string     `json:"mode" xorm:"not null comment('执行方式') VARCHAR(255)"`
	Timeout          int        `json:"timeout" xorm:"not null default 0 comment('超时时间') INT(10)"`
	Retries          int        `json:"retries" xorm:"not null default 0 comment('重试次数') TINYINT(3)"`
	Status           string     `json:"status" xorm:"not null default 'finished' comment('状态') VARCHAR(255)"`
	Result           string     `json:"result" xorm:"not null comment('执行结果') TEXT"`
	Duration         int64      `json:"duration" xorm:"not null comment('持续时间') INT(10)"`
	BeginWith        utils.Time `json:"begin_with" xorm:"not null comment('开始于') DATETIME"`
	FinishWith       utils.Time `json:"finish_with" xorm:"not null comment('结束于') DATETIME"`
	CreatedAt        utils.Time `json:"created_at" xorm:"not null created comment('创建于') DATETIME"`
	UpdatedAt        utils.Time `json:"updated_at" xorm:"not null updated comment('更新于') DATETIME"`
}

// 定义模型的数据表名称
func (records *TaskRecords) TableName() string {
	return "task_records"
}
