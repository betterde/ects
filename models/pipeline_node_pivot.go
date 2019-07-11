package models

type PipelineNodePivot struct {
	PipelineId string `json:"pipeline_id" xorm:"not null pk comment('角色ID') index CHAR(36)"`
	NodeId     string `json:"node_id" xorm:"not null pk comment('节点ID') index CHAR(36)"`
}

// 定义模型的数据表名称
func (pivot *PipelineNodePivot) TableName() string {
	return "pipeline_node_pivot"
}

// Create a pipeline relation
func (pivot *PipelineNodePivot) Store() error {
	_, err := Engine.Insert(pivot)
	return err
}

// Delete a pipeline relation
func (pivot *PipelineNodePivot) Destroy() error {
	_, err := Engine.Delete(pivot)
	return err
}
