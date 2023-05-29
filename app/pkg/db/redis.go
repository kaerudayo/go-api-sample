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
)

var (
	jwtSecret = os.Getenv("JWT_SECRET")
	jwtExpire = time.Hour * time.Duration(8)
)

type TokenCache struct {
	JwtID    string `json:"jwt_id"`
	RecordID string `json:"record_id"`
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
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
	now := time.Now()
	iat := now.Unix()
	exp := now.Add(jwtExpire).Unix()
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
		Iat:      iat,
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
