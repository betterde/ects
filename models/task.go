package models

import (
	"time"
)

const (
	TASK_STATUS_NORMAL   = "normal"
	TASK_STATUS_DISABLED = "disabled"
	TASK_MODE_SHELL      = "shell"
	TASK_MODE_REST       = "rest"
	TASK_MODE_MAIL       = "mail"
	TASK_MODE_HOOK       = "hook"
)

type (
	Task struct {
		ID          string    `json:"id" xorm:"pk char(36) notnull 'id'"`               // ID
		Name        string    `json:"name" xorm:"varchar(255) notnull"`                 // 名称
		ParentID    string    `json:"parent_id" xorm:"char(36) null 'parent_id'"`       // 父任务ID
		Content     string    `json:"content" xorm:"varchar(255) notnull"`              // 任务内容
		Event       string    `json:"event" xorm:"varchar(255)"`                        // 依赖父任务事件
		Mode        string    `json:"mode" xorm:"varchar(255) notnull default 'shell'"` // 执行方式
		Overlap     bool      `json:"overlap" xorm:"tinyint notnull default 0"`         // 重复执行
		Timeout     int       `json:"timeout" xorm:"mediumint notnull default 0"`       // 超时时间
		Interval    int       `json:"interval" xorm:"smallint notnull default 0"`       // 重试间隔
		Retries     int       `json:"retries" xorm:"tinyint notnull default 0"`         // 重试次数
		Status      string    `json:"status" xorm:"char(10) notnull default 'normal'"`  // 状态
		Triggering  time.Time `json:"triggering" xorm:"-"`                              // 下次触发时间
		Description string    `json:"description" xorm:"varchar(255) null"`             // 描述
		CreatedAt   string    `json:"created_at" xorm:"datetime notnull created"`       // 创建时间
		UpdatedAt   string    `json:"updated_at" xorm:"datetime notnull updated"`       // 更新时间
	}
)

// 创建任务
func (task *Task) Store() error {
	_, err := Engine.Insert(task)
	return err
}

// 更新任务
func (task *Task) Update() error {
	_, err := Engine.Id(task.ID).Update(task)
	return err
}

// 删除任务
func (task *Task) Destroy() error {
	_, err := Engine.Delete(task)
	return err
}
