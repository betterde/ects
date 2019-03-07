package models

import (
	"github.com/betterde/ects/internal/server"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	GroupId   int64  `json:"group_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// 为用户颁发Access Token
func (user *User) IssueToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(server.Conf.Auth.TTL * 10e9).Unix()
	claims["uid"] = user.ID
	claims["iat"] = time.Now().Unix()
	claims["issuer"] = "ects"
	claims["username"] = user.Name
	token.Claims = claims

	return token.SignedString([]byte(server.Conf.Auth.Secret))
}
