package models

import (
	"encoding/json"
	"github.com/betterde/ects/internal/utils"
	"github.com/go-xorm/builder"
)

type Pipeline struct {
	Id          string               `json:"id" validate:"-" xorm:"not null pk comment('ID') CHAR(36)"`
	TeamId      string               `json:"team_id" validate:"required,uuid4" xorm:"not null comment('团队ID') index CHAR(36)"`
	Name        string               `json:"name" validate:"required" xorm:"not null comment('名称') VARCHAR(255)"`
	Description string               `json:"description" validate:"-" xorm:"not null comment('描述') VARCHAR(255)"`
	Spec        string               `json:"spec" validate:"required" xorm:"not null comment('定时器') CHAR(64)"`
	Status      int                  `json:"status" validate:"required" xorm:"not null default 1 comment('状态') TINYINT(1)"`
	Finished    string               `json:"finished" validate:"omitempty,uuid4" xorm:"comment('成功时执行') CHAR(36)"`
	Failed      string               `json:"failed" validate:"omitempty,uuid4" xorm:"comment('失败时执行') CHAR(36)"`
	Overlap     int                  `json:"overlap" validate:"numeric" xorm:"not null default 0 comment('重复执行') TINYINT(1)"`
	CreatedAt   utils.Time           `json:"created_at" xorm:"not null created comment('创建于') DATETIME"`
	UpdatedAt   utils.Time           `json:"updated_at" xorm:"not null updated comment('更新于') DATETIME"`
	Nodes       map[string]*Node     `json:"nodes" validate:"-" xorm:"-"`
	Steps       []*PipelineTaskPivot `json:"steps" xorm:"-"`
}

// 定义模型的数据表名称
func (pipeline *Pipeline) TableName() string {
	return "pipelines"
}

// Create a pipeline
func (pipeline *Pipeline) Store() error {
	_, err := Engine.Insert(pipeline)
	return err
}

// Update pipeline attributes
func (pipeline *Pipeline) Update() error {
	_, err := Engine.Id(pipeline.Id).Update(pipeline)
	return err
}

// Delete a pipeline
func (pipeline *Pipeline) Destroy() error {
	_, err := Engine.Delete(pipeline)
	return err
}

// Build pipelne struce to string
func (pipeline *Pipeline) Build() (origin string, err error) {
	relations := make([]*PipelineNodePivot, 0)
	err = Engine.Where(builder.Eq{"pipeline_id": pipeline.Id}).Find(&relations)
	if err != nil {
		return
	}

	ids := make([]string, 0)
	for _, relation := range relations {
		ids = append(ids, relation.NodeId)
	}

	nodes := make(map[string]*Node)

	// Load binded nodes
	err = Engine.Where(builder.Eq{"id": ids}).Find(&nodes)
	if err != nil {
		return
	}

	pipeline.Nodes = nodes

	// Load binded tasks
	err = Engine.Where(builder.Eq{"pipeline_id": pipeline.Id}).Find(&pipeline.Steps)
	if err != nil {
		return
	}

	ob, err := json.Marshal(pipeline)
	origin = string(ob)
	return
}
