package models

import (
	"time"
)

type PasswordResets struct {
	Email     string    `xorm:"not null index VARCHAR(255)"`
	Token     string    `xorm:"not null VARCHAR(255)"`
	CreatedAt time.Time `xorm:"not null created comment('创建于') DATETIME"`
}

// 定义模型的数据表名称
func (resets *PasswordResets) TableName() string {
	return "password_resets"
}
