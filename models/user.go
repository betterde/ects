package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         string `json:"id" xorm:"pk char(36) notnull 'id'"`
	Name       string `json:"name" xorm:"varchar(32) notnull"`
	Email      string `json:"email" xorm:"varchar(64) notnull unique"`
	Password   string `json:"-" xorm:"varchar(128) notnull"`
	Avatar     string `json:"avatar" xorm:"varchar(255) notnull default ''"`
	GroupId    int64  `json:"group_id" xorm:"varchar(36) int index"`
	Manager    bool   `json:"manager" xorm:"tinyint notnull default 0"`
	CreatedAt  string `json:"created_at" xorm:"datetime notnull created"`
	UpdatedAt  string `json:"updated_at" xorm:"datetime null updated"`
	DeletedAt  string `json:"deleted_at" xorm:"datetime null deleted"`
	Model `json:"-" xorm:"-"`
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

func (user *User) Save() error {
	_, err := Engine.Id(user.ID).Update(&user)
	return err
}

func (user *User) ModifyEmail(email string) (*User, error) {
	user.Email = email
	err := user.Save()
	return user, err
}
