package models

// 权限模型
type (
	Permission struct {
		Code        string `json:"code" xorm:"not null pk comment('编码') VARCHAR(255)"`
		Name        string `json:"name" xorm:"not null comment('名称') VARCHAR(255)"`
		Module      string `json:"module" xorm:"not null comment('模块') VARCHAR(255)"`
		Description string `json:"description" xorm:"not null comment('描述') VARCHAR(255)"`
	}
)

// 定义模型的数据表名称
func (permission *Permission) TableName() string {
	return "permissions"
}

// 填充数据
func (permission *Permission) Seed() error {
	return permission.Store()
}

// 创建权限
func (permission *Permission) Store() error {
	_, err := Engine.Insert(permission)
	return err
}
