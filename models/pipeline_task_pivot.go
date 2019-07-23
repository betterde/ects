package models

type PipelineTaskPivot struct {
	PipelineId string `json:"pipeline_id" validate:"required,uuid4" xorm:"not null comment('ID') index CHAR(36)"`
	TaskId     string `json:"task_id" validate:"required,uuid4" xorm:"not null comment('ID') index CHAR(36)"`
	Step       int    `json:"step" validate:"numeric" xorm:"not null comment('步骤') SMALLINT(5)"`
	Timeout    int    `json:"timeout" validate:"numeric" xorm:"not null default 0 comment('超时时间') INT(10)"`
	Interval   int    `json:"interval" validate:"numeric" xorm:"not null default 0 comment('间隔时间') INT(10)"`
	Retries    int    `json:"retries" validate:"numeric" xorm:"not null default 0 comment('重试次数') TINYINT(3)"`
	Dependence string `json:"dependence" validate:"required" xorm:"not null default 'strong' comment('依赖') VARCHAR(255)"`
	CreatedAt  string `json:"-" validate:"-" xorm:"not null created comment('创建于') DATETIME"`
	UpdatedAt  string `json:"-" validate:"-" xorm:"not null updated comment('更新于') DATETIME"`
	Task       *Task  `json:"task" xorm:"-"`
}

// 定义模型的数据表名称
func (pivot *PipelineTaskPivot) TableName() string {
	return "pipeline_task_pivot"
}

// Create a pipeline relation
func (pivot *PipelineTaskPivot) Store() error {
	_, err := Engine.Insert(pivot)
	return err
}

// Delete a pipeline relation
func (pivot *PipelineTaskPivot) Destroy() error {
	_, err := Engine.Delete(pivot)
	return err
}
