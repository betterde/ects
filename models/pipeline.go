package models

import (
	"context"
	"encoding/json"
	"github.com/betterde/ects/internal/utils"
	"github.com/go-xorm/builder"
	"github.com/gorhill/cronexpr"
	uuid "github.com/satori/go.uuid"
	"log"
	"strings"
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
	FinishedTask *Task                `json:"finished_task" xorm:"-"`
	FailedTask   *Task                `json:"failed_task" xorm:"-"`
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

// 执行流水线
func (pipeline *Pipeline) Exec(ctx context.Context) {
	if len(pipeline.Steps) > 0 {
		record := &PipelineRecords{
			Id:         uuid.NewV4().String(),
			PipelineId: pipeline.Id,
			NodeId:     "",
			WorkerName: "",
			Spec:       pipeline.Spec,
			Status:     1,
			Duration:   0,
		}
		beginWith := time.Now()
		// 按照任务的排序，逐个执行
		for _, pivot := range pipeline.Steps {
			var pctx context.Context
			var cancelFunc func()

			// 如果设置了超时时间则
			if pivot.Timeout == 0 {
				pctx, cancelFunc = context.WithCancel(ctx)
			} else {
				pctx, cancelFunc = context.WithTimeout(ctx, time.Duration(pivot.Timeout)*time.Second)
			}
			env := strings.Split(pivot.Environment, " ")
			taskRecord, err := pivot.Task.Exec(pctx, pivot.User, pivot.Directory, env)
			if err != nil {
				record.Status = 0
				goto END
			}

			taskRecord.NodeId = ""
			taskRecord.Timeout = pivot.Timeout
			taskRecord.Retries = pivot.Retries
			taskRecord.WorkerName = ""
			taskRecord.CreatedAt = utils.Time(time.Now())
			taskRecord.UpdatedAt = utils.Time(time.Now())

			select {
			case <-ctx.Done():
				cancelFunc()
				break
			default:
				continue
			}
		}
	END:
		finishWith := time.Now()
		record.Duration = int64(finishWith.Sub(beginWith).Seconds())
		record.BeginWith = utils.Time(beginWith)
		record.FinishWith = utils.Time(finishWith)
		record.FinishWith = utils.Time(time.Now())
		record.CreatedAt = utils.Time(time.Now())
		record.UpdatedAt = utils.Time(time.Now())
		if err := record.Store(); err != nil {
			// TODO
		}

		// 当流水线成功时触发
		if record.Status == 1 && pipeline.Finished != "" {
			if _, err := pipeline.FinishedTask.Exec(context.Background(), "", "", []string{}); err != nil {
				log.Println(err)
			}
		}

		// 当流水线失败时触发
		if record.Status == 0 && pipeline.Failed != "" {
			if _, err := pipeline.FailedTask.Exec(context.Background(), "", "", []string{}); err != nil {
				log.Println(err)
			}
		}
	}

	return
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

	finishedTask := tasks[pipeline.Finished]
	failedTask := tasks[pipeline.Failed]
	pipeline.FinishedTask = &finishedTask
	pipeline.FailedTask = &failedTask

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
