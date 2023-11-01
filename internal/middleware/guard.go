package middleware

import (
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
)

type Claims struct {
	Foo string `json:"foo"`
}

var (
	secret = []byte("signature_hmac_secret_shared_key")

	JWTHandler context.Handler
)

func init() {
	verifier := jwt.NewVerifier(jwt.HS256, secret)
	verifier.WithDefaultBlocklist()

	JWTHandler = verifier.Verify(func() interface{} {
		return new(Claims)
	})
}
