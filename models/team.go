package models

type (
	Team struct {
		ID          string  `json:"id" xorm:"pk char(36) notnull 'id'"`
		Name        string  `json:"name" xorm:"varchar(36) notnull"`
		Avatar      string  `json:"avatar" xorm:"varchar(255) null"`
		Description string  `json:"description" xorm:"varchar(255) null"`
		Users       []*User `json:"users" xorm:"-"`
	}
)
