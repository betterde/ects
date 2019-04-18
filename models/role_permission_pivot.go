package models

type RolePermissionPivot struct {
	RoleId         string `json:"role_id" validate:"required,uuid4" xorm:"not null pk comment('角色ID') index CHAR(36)"`
	PermissionCode string `json:"permission_code" validate:"required" xorm:"not null pk comment('权限编码') index VARCHAR(255)"`
}

// 定义模型的数据表名称
func (pivot *RolePermissionPivot) TableName() string {
	return "role_permission_pivot"
}