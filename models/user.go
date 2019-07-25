package models

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id        string    `json:"id" xorm:"not null pk comment('用户ID') CHAR(36)"`
	Name      string    `json:"name" xorm:"not null comment('姓名') VARCHAR(255)"`
	Email     string    `json:"email" xorm:"not null comment('邮箱') unique VARCHAR(255)"`
	Password  string    `json:"-" xorm:"not null comment('密码') VARCHAR(255)"`
	Avatar    string    `json:"avatar" xorm:"null comment('头像') VARCHAR(255)"`
	TeamId    string    `json:"team_id" xorm:"not null default '' comment('团队ID') index CHAR(36)"`
	RoleId    string    `json:"role_id" xorm:"not null default '' comment('角色ID') index CHAR(36)"`
	Manager   bool      `json:"manager" xorm:"not null default 0 comment('管理员') TINYINT(1)"`
	CreatedAt time.Time `json:"created_at" xorm:"not null created comment('创建于') DATETIME"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null updated comment('更新于') DATETIME"`
	DeletedAt time.Time `json:"deleted_at" xorm:"null comment('删除于') DATETIME"`
}

// Define table name
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

func (user *User) MarshalJSON() ([]byte, error) {
	type Alias User

	var DeletedAt = ""

	if !user.DeletedAt.IsZero() {
		DeletedAt = user.DeletedAt.Format(DefaultTimeFormat)
	}

	return json.Marshal(&struct {
		Alias
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		DeletedAt string `json:"deleted_at"`
	}{
		Alias:     Alias(*user),
		CreatedAt: user.CreatedAt.Format(DefaultTimeFormat),
		UpdatedAt: user.UpdatedAt.Format(DefaultTimeFormat),
		DeletedAt: DeletedAt,
	})
}
