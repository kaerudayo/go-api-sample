package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/api-sample/app/infra"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

var (
	jwtSecret = os.Getenv("JWT_SECRET")
	jwtExpire = time.Minute * time.Duration(30)
)

type TokenCache struct {
	JwtID    string `json:"jwt_id"`
	RecordID string `json:"record_id"`
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
}

// jwt tokenの検証
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
			if err.Error() != "Token used before issued" {
				if err != nil {
					return nil, err
				}
				if !token.Valid {
					return nil, errors.New("invalid token")
				}
			}

			return token, nil
		},
	})
}

// contextからjwt tokenを取得
func GetToken(c echo.Context) (TokenCache, error) {
	rawToken := c.Get("rawToken").(*jwt.Token)
	claims := rawToken.Claims.(jwt.MapClaims)
	if _, ok := claims["id"]; !ok {
		return TokenCache{}, errors.New("claims is empty")
	}

	jwtID := claims["id"]
	key := fmt.Sprintf("%s/%s", jwtSecret, jwtID.(string))
	val, err := infra.RedisDB.Get(infra.Ctx, key).Result()
	if val != "" && err != redis.Nil {
		data := TokenCache{}
		err := json.Unmarshal([]byte(val), &data)
		if err != nil {
			return TokenCache{}, err
		}
		return data, nil
	}

	return TokenCache{}, errors.New("value not found")
}

func CreateAndSetToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	iat := time.Now()
	exp := iat.Add(jwtExpire).Unix()
	jwtID := GenID()

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = jwtID
	claims["record_id"] = id
	claims["iat"] = iat
	claims["exp"] = exp

	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	data, err := json.Marshal(TokenCache{
		JwtID:    jwtID,
		RecordID: id,
		Iat:      iat.Unix(),
		Exp:      exp,
	})
	if err != nil {
		return "", err
	}

	key := fmt.Sprintf("%s/%s", infra.CacheJwtTokens, jwtID)
	err = infra.RedisDB.Set(infra.Ctx, key, data, jwtExpire).Err()
	if err != nil {
		return "", err
	}

	return t, nil
}
