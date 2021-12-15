package redis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
	"time"
)

var RedisClient *redis.Client

type RedisConfig struct {
	Address  string `env:""`
	Pwd      string `env:""`
	DB       int    `env:""`
	UserName string `env:""`
	PoolSize int    `env:""`
}

func init() {
	redisConfig := RedisConfig{
		Address:  "127.0.0.1:6379",
		PoolSize: 200,
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Pwd,
		DB:       redisConfig.DB,
		PoolSize: redisConfig.PoolSize,
	})
}

const (
	ObjUrlKey = "CityIot:ObjUrlKey:%d"
)

func GetCache(key string) (string, error) {
	result, err := RedisClient.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			logrus.Warnf("未找到缓存数据，key:%s", key)
			return "", nil
		}
		logrus.Errorf("查询缓存失败:%v, key:%s", err, key)
		if err != nil {
			return "", fmt.Errorf("查询缓存失败")
		}
	}
	return result, nil
}

func SetCache(key, val string, exp time.Duration) error {
	_, err := RedisClient.SetNX(key, val, exp).Result()
	if err != nil {
		logrus.Errorf("缓存写入失败:%v, key:%s", err, key)
		if err != nil {
			return fmt.Errorf("缓存写入失败")
		}
	}
	return nil
}
