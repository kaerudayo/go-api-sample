package infra

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/api-sample/app/pkg/consts"
	"github.com/go-redis/redis/v8"
)

var (
	redisHost                = os.Getenv("REDIS_HOST")
	redisPort                = os.Getenv("REDIS_PORT")
	redisPwd                 = os.Getenv("REDIS_PWD")
	ctx                      = context.Background()
	redisCachExpier          = time.Duration(24) * time.Hour
	CacheLoginToken          = consts.Env + "/token"
	CacheRequestLimitWatcher = consts.Env + "/requestLimitWatcher"
	CacheUsers               = consts.Env + "/users"
	CacheJwtTokens           = consts.Env + "/jwt"
	CacheProposals           = consts.Env + "/proposals"
	CachePwdResetToken       = consts.Env + "/password"
	CacheDB                  = consts.Env + "/db"
	RedisDB                  *redis.Client
	Ctx                      = context.Background()
)

func CashInit() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPwd,
		DB:       0,
	})
}

func Get(key string) string {
	val, err := RedisDB.Get(ctx, key).Result()
	if err != redis.Nil {
		return val
	}
	return ""
}

func Set(key string, data string, expire *time.Duration) bool {
	if expire == nil {
		expire = &redisCachExpier
	}

	err := RedisDB.Set(ctx, key, data, *expire).Err()
	if err != nil {
		panic(err)
	}
	return true
}

func Keys(pattern string) []string {
	val, err := RedisDB.Keys(ctx, pattern).Result()
	if err != redis.Nil {
		return val
	}

	return []string{}
}

func Del(key string) bool {
	err := RedisDB.Del(ctx, key).Err()
	if err != nil {
		panic(err)
	}

	return true
}

func ClearTestData() {
	var cursor uint64
	keys, _, err := RedisDB.Scan(ctx, cursor, "test/*", 0).Result()
	if err == nil {
		for _, key := range keys {
			RedisDB.Del(ctx, key)
		}
	}
}
