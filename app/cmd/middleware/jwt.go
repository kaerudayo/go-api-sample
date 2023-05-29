package middleware

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

var (
	jwtSecret []byte = []byte(os.Getenv("JWT_SECRET"))
)

// jwt tokenの検証
// https://jwt.io/
func ValidateToken() echo.MiddlewareFunc {
	return m.JWTWithConfig(m.JWTConfig{
		SigningKey: jwtSecret,
		ContextKey: "rawToken",
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				return jwtSecret, nil
			}

			token, err := jwt.Parse(auth, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			return token, nil
		},
	})
}
