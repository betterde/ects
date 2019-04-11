package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id        string    `xorm:"not null pk comment('用户ID') CHAR(36)"`
	Name      string    `xorm:"not null comment('姓名') VARCHAR(255)"`
	Email     string    `xorm:"not null comment('邮箱') unique VARCHAR(255)"`
	Password  string    `xorm:"not null comment('密码') VARCHAR(255)"`
	Avatar    string    `xorm:"comment('头像') VARCHAR(255)"`
	TeamId    string    `xorm:"not null default '' comment('团队ID') index CHAR(36)"`
	RoleId    string    `xorm:"not null default '' comment('角色ID') index CHAR(36)"`
	Manager   bool      `xorm:"not null default 0 comment('管理员') TINYINT(1)"`
	CreatedAt time.Time `xorm:"not null created comment('创建于') DATETIME"`
	UpdatedAt time.Time `xorm:"not null updated comment('更新于') DATETIME"`
	DeletedAt time.Time `xorm:"null comment('删除于') DATETIME"`
}

// 定义模型的数据表名称
func (user *User) TableName() string {
	return "users"
}

func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ValidatePassword(password string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(password)); err != nil {
		return false, err
	}
	return true, nil
}

func (user *User) Store() error {
	_, err := Engine.Insert(user)
	return err
}

func (user *User) Save() error {
	_, err := Engine.Id(user.Id).Update(&user)
	return err
}

func (user *User) ModifyEmail(email string) (*User, error) {
	user.Email = email
	err := user.Save()
	return user, err
}
