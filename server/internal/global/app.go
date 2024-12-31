package global

import (
	"simple-tool/server/config/autoload"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	// BasePath 项目根目录
	BasePath string
	// Conf 全局配置
	Conf *autoload.AppConfig
	// ZapLog 全局日志指针
	ZapLog *zap.Logger
	// DB 全局Db
	DB *gorm.DB
	// RDb 全局redis
	RDb *redis.Client
)
