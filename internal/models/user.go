package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string `json:"id" xorm:"char(32) notnull unique 'id'"`
	Name      string `json:"name" xorm:"varchar(32) notnull"`
	Email     string `json:"email" xorm:"varchar(64) notnull unique"`
	Password  string `json:"-" xorm:"varchar(128) notnull"`
	Avatar    string `json:"avatar" xorm:"varchar(255) null"`
	GroupId   int64  `json:"group_id" xorm:"varchar(32) int index"`
	CreatedAt string `json:"created_at" xorm:"datetime notnull created"`
	UpdatedAt string `json:"updated_at" xorm:"datetime null updated"`
	Model     `json:"-" xorm:"-"`
}

func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(userPassword)); err != nil {
		return false, err
	}
	return true, nil
}

func (user *User) Store() *User {
	return user
}
