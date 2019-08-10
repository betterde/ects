package models

import "github.com/betterde/ects/internal/utils"

type PipelineNodePivot struct {
	PipelineId string     `json:"pipeline_id" validate:"required,uuid4" xorm:"not null pk comment('角色ID') CHAR(36)"`
	NodeId     string     `json:"node_id" validate:"required,uuid4" xorm:"not null pk comment('节点ID') CHAR(36)"`
	CreatedAt  utils.Time `json:"created_at" validate:"-" xorm:"not null created comment('创建于') DATETIME"`
	Pipeline   *Pipeline  `json:"pipeline" xorm:"-"`
}

// 定义模型的数据表名称
func (pivot *PipelineNodePivot) TableName() string {
	return "pipeline_node_pivot"
}

// 创建关联
func (pivot *PipelineNodePivot) Store() error {
	_, err := Engine.Insert(pivot)
	return err
}

// 解除关联
func (pivot *PipelineNodePivot) Destroy() error {
	_, err := Engine.Delete(pivot)
	return err
}
