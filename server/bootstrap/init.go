package bootstrap

import (
	"simple-tool/server/internal/global"
	"simple-tool/server/pkg/config"
	"simple-tool/server/pkg/db"
	"simple-tool/server/pkg/logger"
	"simple-tool/server/pkg/redis"
	"simple-tool/server/pkg/validator"
)

func init() {
	// 加载配置
	config.Init()
	// 初始化日志
	global.ZapLog = logger.Init()
	// 初始化mysql
	global.DB = db.Init()
	// 初始化redis
	global.RDb = redis.Init()
	// 初始化验证器
	validator.Init()
}
