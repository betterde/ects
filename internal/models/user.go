package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/betterde/ects/internal/utils/random"
)

type User struct {
	ID        string `json:"id" xorm:"char(32) notnull unique 'id'"`
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

// 生成密码
func (user *User) PasswordGenerator(origin string) {
	user.Salt = random.String(16)
	m := md5.New()
	m.Write([]byte(origin))
	user.Password = hex.EncodeToString(m.Sum(nil))
}

func (user *User) Store() *User {
	return user
}
