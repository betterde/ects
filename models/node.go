package models

import (
	"github.com/go-xorm/builder"
)

const (
	ONLINE      = "online"
	OFFLINE     = "offline"
	MODE_MASTER = "master"
	MODE_WORKER = "worker"
)

type (
	Node struct {
		Id          string `json:"id" xorm:"not null pk comment('用户ID') CHAR(36)"`
		Name        string `json:"name" xorm:"not null comment('名称') VARCHAR(255)"`
		Host        string `json:"host" xorm:"not null comment('主机地址') VARCHAR(255)"`
		Port        int    `json:"port" xorm:"not null comment('端口') SMALLINT(5)"`
		Mode        string `json:"mode" xorm:"not null comment('节点类型') CHAR(6)"`
		Status      string `json:"status" xorm:"not null default('connected') comment('状态') VARCHAR(255)"`
		Description string `json:"description" xorm:"comment('描述') VARCHAR(255)"`
		CreatedAt   string `json:"created_at" xorm:"not null created comment('创建于') DATETIME"`
		UpdatedAt   string `json:"updated_at" xorm:"not null updated comment('更新于') DATETIME"`
	}
)

// 定义模型的数据表名称
func (node *Node) TableName() string {
	return "nodes"
}

// 创建节点
func (node *Node) Store() error {
	_, err := Engine.Insert(node)
	return err
}

// 更新节点信息
func (node *Node) Update() error {
	_, err := Engine.Id(node.Id).Update(node)
	return err
}

// 更新状态为在线
func (node *Node) Online() {
	node.Status = ONLINE
	if err := node.Update(); err != nil {
		// todo
	}
}

// 更新状态为离线
func (node *Node) Offline() {
	node.Status = OFFLINE
	if err := node.Update(); err != nil {
		// todo
	}
}

// 创建或更新节点
func (node *Node) CreateOrUpdate() error {
	exist := &Node{}
	if count, err := Engine.Where(builder.Eq{"id": node.Id}).Count(exist); count > 0 && err == nil {
		return node.Update()
	}

	return node.Store()
}
