package models

import "time"

type (
	Role struct {
		Id          string    `xorm:"not null pk comment('ID') CHAR(36)"`
		Name        string    `xorm:"not null comment('名称') VARCHAR(255)"`
		Description string    `xorm:"not null comment('描述') VARCHAR(255)"`
		System      int       `xorm:"not null default 0 comment('系统定义') TINYINT(1)"`
		CreatedAt   time.Time `xorm:"not null created comment('创建于') DATETIME"`
		UpdatedAt   time.Time `xorm:"not null updated comment('更新于') DATETIME"`
	}
)

// 定义模型的数据表名称
func (role *Role) TableName() string {
	return "roles"
}

// 数据填充器
func (role *Role) Seed() error {
	return role.Store()
}

// 创建角色
func (role *Role) Store() error {
	_, err := Engine.Insert(role)
	return err
}