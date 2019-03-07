package models

type User struct {
	ID        int64  `json:"id" xorm:"char(32) notnull unique 'id'"`
	Name      string `json:"name" xorm:"varchar(32) notnull"`
	Email     string `json:"email" xorm:"varchar(64) notnull unique"`
	Salt      string `json:"-" xorm:"char(16) notnull"`
	Password  string `json:"-" xorm:"varchar(128) notnull"`
	Avatar    string `json:"avatar" xorm:"varchar(255) null"`
	GroupId   int64  `json:"group_id" xorm:"varchar(32) int index"`
	CreatedAt string `json:"created_at" xorm:"datetime notnull created"`
	UpdatedAt string `json:"updated_at" xorm:"datetime null updated"`
	Model     `json:"-" xorm:"-"`
}

func (user *User) SetPassword(origin, salt string) {
	user.Salt = salt

	user.Password = ""
	return
}