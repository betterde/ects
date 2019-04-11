package models

type PipelineNodePivot struct {
	RoleId string `xorm:"not null pk comment('角色ID') index CHAR(36)"`
	NodeId string `xorm:"not null pk comment('节点ID') index CHAR(36)"`
}

// 定义模型的数据表名称
func (pivot *PipelineNodePivot) TableName() string {
	return "pipeline_node_pivot"
}
