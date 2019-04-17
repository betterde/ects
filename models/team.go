package models

import "time"

type (
	Team struct {
		Id          string    `json:"id" validate:"-" xorm:"not null pk comment('ID') CHAR(36)"`
		Name        string    `json:"name" validate:"required" xorm:"not null comment('名称') VARCHAR(255)"`
		Description string    `json:"description" validate:"-" xorm:"not null comment('描述') VARCHAR(255)"`
		CreatedAt   time.Time `json:"created_at" validate:"-" xorm:"not null created comment('创建于') DATETIME"`
		UpdatedAt   time.Time `json:"updated_at" validate:"-" xorm:"not null updated comment('更新于') DATETIME"`
	}
)

// 定义模型的数据表名称
func (team *Team) TableName() string {
	return "teams"
}

func (team *Team) Store() error {
	_, err := Engine.Insert(team)
	return err
}

func (team *Team) Update() error {
	_, err := Engine.Id(team.Id).Update(team)
	return err
}

func (team *Team) Destroy() error {
	_, err := Engine.Delete(team)
	return err
}