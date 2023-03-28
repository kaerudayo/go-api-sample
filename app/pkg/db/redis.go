package db

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/api-sample/app/infra"
	"github.com/api-sample/app/pkg/logger"
	"github.com/golang-jwt/jwt/v4"
)

var (
	jwtSecret = os.Getenv("JWT_SECRET")
	jwtExpire = time.Minute * time.Duration(15)
)

type TokenCache struct {
	JwtID    string `json:"jwt_id"`
	ID       string `json:"id"`
	RecordID string `json:"record_id"`
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
}

func DBKeyGen(tableName string, args []string) string {
	if args[0] == "all" {
		return fmt.Sprintf("%s/%s/*", infra.CacheDB, tableName)
	}
	res, err := json.Marshal(args)
	if err != nil {
		logger.SugerError(err.Error())
	}
	return fmt.Sprintf("%s/%s/%s", infra.CacheDB, tableName, string(res))
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
