package models

import "time"

const (
	STATUS_CONNECTED    = "connected"
	STATUS_DISCONNECTED = "disconnected"
)

type (
	Worker struct {
		Id          string    `xorm:"not null pk comment('用户ID') CHAR(36)"`
		Name        string    `xorm:"not null comment('名称') VARCHAR(255)"`
		Host        string    `xorm:"not null comment('主机地址') VARCHAR(255)"`
		Port        int       `xorm:"not null comment('端口') SMALLINT(5)"`
		Status      string    `xorm:"not null default('connected') comment('状态') VARCHAR(255)"`
		Description string    `xorm:"comment('描述') VARCHAR(255)"`
		CreatedAt   time.Time `xorm:"not null created comment('创建于') TIMESTAMP"`
		UpdatedAt   time.Time `xorm:"not null updated comment('更新于') TIMESTAMP"`
	}
)

// 定义模型的数据表名称
func (worker *Worker) TableName() string {
	return "workers"
}

// 创建节点
func (worker *Worker) Store() error {
	_, err := Engine.Insert(worker)
	return err
}

// 更新节点信息
func (worker *Worker) Update() error {
	_, err := Engine.Id(worker.Id).Update(worker)
	return err
}
