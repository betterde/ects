package models

import (
	"time"
)

type Pipeline struct {
	Id          string    `json:"id" validate:"-" xorm:"not null pk comment('ID') CHAR(36)"`
	TeamId      string    `json:"team_id" validate:"required,uuid4" xorm:"not null comment('团队ID') index CHAR(36)"`
	Name        string    `json:"name" validate:"required" xorm:"not null comment('名称') VARCHAR(255)"`
	Description string    `json:"description" validate:"-" xorm:"not null comment('描述') VARCHAR(255)"`
	Spec        string    `json:"spec" validate:"required" xorm:"not null comment('定时器') CHAR(64)"`
	Status      int       `json:"status" validate:"required" xorm:"not null default 1 comment('状态') TINYINT(1)"`
	Finished    string    `json:"finished" validate:"omitempty,uuid4" xorm:"comment('成功时执行') CHAR(36)"`
	Failed      string    `json:"failed" validate:"omitempty,uuid4" xorm:"comment('失败时执行') CHAR(36)"`
	Overlap     int       `json:"overlap" validate:"numeric" xorm:"not null default 0 comment('重复执行') TINYINT(1)"`
	CreatedAt   time.Time `json:"created_at" xorm:"not null created comment('创建于') DATETIME"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"not null updated comment('更新于') DATETIME"`
}

// 定义模型的数据表名称
func (pipelines *Pipeline) TableName() string {
	return "pipelines"
}

// Create a pipeline
func (pipelines *Pipeline) Store() error {
	_, err := Engine.Insert(pipelines)
	return err
}

// Update pipeline attributes
func (pipelines *Pipeline) Update() error {
	_, err := Engine.Id(pipelines.Id).Update(pipelines)
	return err
}

// Delete a pipeline
func (pipelines *Pipeline) Destroy() error {
	_, err := Engine.Delete(pipelines)
	return err
}
