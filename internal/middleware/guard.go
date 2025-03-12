package middleware

import (
	"github.com/betterde/ects/config"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
)

var (
	JWTHandler context.Handler
)

func Init() context.Handler {
	secret := []byte(config.Conf.Auth.Secret)
	verifier := jwt.NewVerifier(jwt.HS256, secret)
	verifier.WithDefaultBlocklist()

	if JWTHandler != nil {
		return JWTHandler
	}

	JWTHandler = verifier.Verify(func() interface{} {
		return new(jwt.Claims)
	})

	return JWTHandler
}
