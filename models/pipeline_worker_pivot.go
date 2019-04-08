package models

type PipelineWorkerPivot struct {
	RoleId   string `xorm:"not null pk comment('角色ID') index CHAR(36)"`
	WorkerId string `xorm:"not null pk comment('节点ID') index CHAR(36)"`
}

// 定义模型的数据表名称
func (pivot *PipelineWorkerPivot) TableName() string {
	return "pipeline_worker_pivot"
}
