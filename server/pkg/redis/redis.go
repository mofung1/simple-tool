package redis

import (
	"fmt"
	"log"
	"simple-tool/server/internal/global"
	"time"

	"github.com/go-redis/redis"
)

// Init 初始化连接
func Init() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.Conf.RedisConfig.Host, global.Conf.RedisConfig.Port),
		Password: global.Conf.RedisConfig.Password,     // no password set
		DB:       global.Conf.RedisConfig.DB,           // use default DB
		PoolSize: global.Conf.RedisConfig.MinIdleConns, // 连接池大小
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		global.ZapLog.Error("redis bootstrap failed,err:" + err.Error())
		log.Fatal("redis连接失败" + err.Error())
	}

	return rdb
}

// Set 设置缓存
func Set(key string, value interface{}, t int64) bool {
	expire := time.Duration(t) * time.Second
	key = GetFullKey(key)
	if err := global.RDb.Set(key, value, expire).Err(); err != nil {
		return false
	}
	return true
}

// GetString 获取字符串缓存
func GetString(key string) string {
	key = GetFullKey(key)
	result, err := global.RDb.Get(key).Result()
	if err != nil {
		return ""
	}
	return result
}

// GetInt 获取int类型缓存
func GetInt(key string) int {
	key = GetFullKey(key)
	result, err := global.RDb.Get(key).Int()
	if err != nil {
		return 0
	}
	return result
}

// GetInt64 获取int64类型缓存
func GetInt64(key string) int64 {
	key = GetFullKey(key)
	result, err := global.RDb.Get(key).Int64()
	if err != nil {
		return 0
	}
	return result
}

// Del 删除缓存
func Del(key string) bool {
	key = GetFullKey(key)
	_, err := global.RDb.Del(key).Result()
	if err != nil {
		global.ZapLog.Error("redis-del失败:" + err.Error())
		return false
	}
	return true
}

// Expire 设置过期时间
func Expire(key string, t int64) bool {
	// 延长过期时间
	key = GetFullKey(key)
	expire := time.Duration(t) * time.Second
	if err := global.RDb.Expire(key, expire).Err(); err != nil {
		global.ZapLog.Error("redis-expire失败:" + err.Error())
		return false
	}
	return true
}

// GetFullKey 获取完整前缀key
func GetFullKey(key string) string {
	return global.Conf.Name + ":" + key
}
