package models

type Configuration struct {
	Name        string `xorm:"not null pk comment('名称') VARCHAR(255)"`
	Value       string `xorm:"not null comment('当前值') VARCHAR(255)"`
	Default     string `xorm:"not null comment('默认值') VARCHAR(255)"`
	Module      string `xorm:"not null comment('作用范围') VARCHAR(255)"`
	Description string `xorm:"not null comment('描述') VARCHAR(255)"`
}

// 定义模型的数据表名称
func (configuration *Configuration) TableName() string {
	return "configuration"
}
