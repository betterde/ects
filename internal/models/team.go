package models

type (
	Team struct {
		ID string `json:"id" xorm:"pk char(36) notnull 'id'"`
		Name string `json:"name" xorm:"varchar(36) notnull 'name'"`
		Avatar string `json:"avatar" xorm:"varchar(255) null 'avatar'"`
	}
)