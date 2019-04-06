package models

type (
	Team struct {
		ID          string  `json:"id" xorm:"pk char(36) notnull 'id'"`
		Name        string  `json:"name" xorm:"varchar(36) notnull"`
		Description string  `json:"description" xorm:"varchar(255) null"`
		CreatedAt   string  `json:"created_at" xorm:"datetime notnull created"`
		UpdatedAt   string  `json:"updated_at" xorm:"datetime null updated"`
		DeletedAt   string  `json:"deleted_at" xorm:"datetime null deleted"`
		Users       []*User `json:"users" xorm:"-"`
	}
)
