package configuration

import (
	"context"
	"encoding/json"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
	"myponyasia.com/hub-api/exception"
)

var RDB *redis.Client

func RedisConfig() {
	REDIS_HOST := os.Getenv("REDIS_HOST")
	REDIS_PORT := os.Getenv("REDIS_PORT")
	REDIS_USER := os.Getenv("REDIS_USER")
	REDIS_PASS := os.Getenv("REDIS_PASS")
	maxPoolSize, err := strconv.Atoi(os.Getenv("REDIS_POOL_MAX_SIZE"))
	exception.PanicLogging(err)
	minIdlePoolSize, err := strconv.Atoi(os.Getenv("REDIS_POOL_MIN_IDLE_SIZE"))
	exception.PanicLogging(err)

	RDB = redis.NewClient(&redis.Options{
		Addr:         REDIS_HOST + ":" + REDIS_PORT,
		Username:     REDIS_USER,
		Password:     REDIS_PASS,
		DB:           0,
		PoolSize:     maxPoolSize,
		MinIdleConns: minIdlePoolSize,
	})
}

func SetCache[T any](ctx context.Context, prefix string, key string, executeData func(context.Context, string) (T, error)) *T {
	var data []byte
	var object T
	if err := RDB.Get(ctx, prefix+"_"+key).Scan(&data); err == nil {
		err := json.Unmarshal(data, &object)
		exception.PanicLogging(err)

		return &object
	}
	value, err := executeData(ctx, key)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	cacheValue, err := json.Marshal(value)
	exception.PanicLogging(err)

	if err := RDB.Set(ctx, prefix+"_"+key, cacheValue, -1).Err(); err != nil {
		exception.PanicLogging(err)
	}
	return &value
}
