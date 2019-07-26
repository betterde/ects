package models

import "github.com/betterde/ects/internal/utils"

type (
	Role struct {
		Id          string     `json:"id" validate:"-" xorm:"not null pk comment('ID') CHAR(36)"`
		Name        string     `json:"name" validate:"required" xorm:"not null comment('名称') VARCHAR(255)"`
		Description string     `json:"description" validate:"-" xorm:"not null comment('描述') VARCHAR(255)"`
		System      int        `json:"system" validate:"-" xorm:"not null default 0 comment('系统定义') TINYINT(1)"`
		CreatedAt   utils.Time `json:"created_at" validate:"-" xorm:"not null created comment('创建于') DATETIME"`
		UpdatedAt   utils.Time `json:"updated_at" validate:"-" xorm:"not null updated comment('更新于') DATETIME"`
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

// Create a pipeline
func (role *Role) Store() error {
	_, err := Engine.Insert(role)
	return err
}

// Update role attributes
func (role *Role) Update() error {
	_, err := Engine.Id(role.Id).Update(role)
	return err
}

// Delete a pipeline
func (role *Role) Destroy() error {
	_, err := Engine.Delete(role)
	return err
}
