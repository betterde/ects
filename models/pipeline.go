package models

import (
	"encoding/json"
	"github.com/betterde/ects/internal/utils"
	"github.com/go-xorm/builder"
	"github.com/gorhill/cronexpr"
	"time"
)

// 流水线模型
type Pipeline struct {
	Id           string               `json:"id" validate:"-" xorm:"not null pk comment('ID') CHAR(36)"`
	Name         string               `json:"name" validate:"required" xorm:"not null comment('名称') VARCHAR(255)"`
	Description  string               `json:"description" validate:"-" xorm:"not null comment('描述') VARCHAR(255)"`
	Spec         string               `json:"spec" validate:"required" xorm:"not null comment('定时器') CHAR(64)"`
	Status       int                  `json:"status" validate:"numeric" xorm:"not null default 0 comment('状态') TINYINT(1)"`
	Finished     string               `json:"finished" validate:"omitempty,uuid4" xorm:"null comment('成功时执行') CHAR(36)"`
	Failed       string               `json:"failed" validate:"omitempty,uuid4" xorm:"null comment('失败时执行') CHAR(36)"`
	Overlap      int                  `json:"overlap" validate:"numeric" xorm:"not null default 0 comment('重复执行') TINYINT(1)"`
	CreatedAt    utils.Time           `json:"created_at" validate:"-" xorm:"not null created comment('创建于') DATETIME"`
	UpdatedAt    utils.Time           `json:"updated_at" validate:"-" xorm:"not null updated comment('更新于') DATETIME"`
	Nodes        []string             `json:"nodes" xorm:"-"`
	Steps        []*PipelineTaskPivot `json:"steps" xorm:"-"`
	Expression   *cronexpr.Expression `json:"-" xorm:"-"`
	NextTime     time.Time            `json:"-" xorm:"-"`
	FinishedTask *Task                `json:"finished_task,omitempty" xorm:"-"`
	FailedTask   *Task                `json:"failed_task,omitempty" xorm:"-"`
}

// 定义模型的数据表名称
func (pipeline *Pipeline) TableName() string {
	return "pipelines"
}

// 创建任务流水线
func (pipeline *Pipeline) Store() error {
	_, err := Engine.Insert(pipeline)
	return err
}

// 更新任务流水线属性
func (pipeline *Pipeline) Update() error {
	_, err := Engine.Id(pipeline.Id).Update(pipeline)
	return err
}

// 删除任务流水线
func (pipeline *Pipeline) Destroy() error {
	_, err := Engine.Delete(pipeline)
	return err
}

// 构造流水线数据结构
func (pipeline *Pipeline) Build() (origin []byte, err error) {
	relations := make([]*PipelineNodePivot, 0)

	if err = Engine.Where(builder.Eq{"pipeline_id": pipeline.Id}).Find(&relations); err != nil {
		return []byte{}, err
	}

	for _, relation := range relations {
		pipeline.Nodes = append(pipeline.Nodes, relation.NodeId)
	}

	if err = Engine.Where(builder.Eq{"pipeline_id": pipeline.Id}).Find(&pipeline.Steps); err != nil {
		return []byte{}, err
	}

	ids := make([]string, 0)
	for _, relation := range pipeline.Steps {
		ids = append(ids, relation.TaskId)
	}

	ids = append(ids, pipeline.Finished)
	ids = append(ids, pipeline.Failed)

	tasks := make(map[string]Task)
	if err := Engine.Where(builder.Eq{"id": ids}).Find(&tasks); err != nil {
		return []byte{}, err
	}

	finishedTask, exist := tasks[pipeline.Finished]
	if exist {
		pipeline.FinishedTask = &finishedTask
	}
	failedTask, exist := tasks[pipeline.Failed]
	if exist {
		pipeline.FailedTask = &failedTask
	}

	for index, relation := range pipeline.Steps {
		task := tasks[relation.TaskId]
		pipeline.Steps[index].Task = &task
	}

	origin, err = json.Marshal(pipeline)
	return
}

// 序列化
func (pipeline *Pipeline) ToString() (string, error) {
	result, err := json.Marshal(pipeline)
	return string(result), err
}
