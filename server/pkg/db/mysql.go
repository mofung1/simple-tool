package db

import (
	"fmt"
	"io"
	"log"
	"os"
	"simple-tool/server/internal/global"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Init 初始化数据库
func Init() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=true&loc=Local",
		global.Conf.DBConfig.User,
		global.Conf.DBConfig.Password,
		global.Conf.DBConfig.Host,
		global.Conf.DBConfig.DBName,
		global.Conf.DBConfig.Charset,
	)

	configs := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,                        // 使用单数表名
			TablePrefix:   global.Conf.DBConfig.Prefix, // 表名前缀
		},
		Logger:                 getGormLogger(),
		SkipDefaultTransaction: true,
	}

	gormDB, err := gorm.Open(mysql.Open(dsn), configs)

	if err != nil {
		global.ZapLog.Error("connect DB failed", zap.Error(err))
		log.Fatal("mysql连接失败" + err.Error())
	}

	db, err := gormDB.DB()
	if err != nil {
		global.ZapLog.Error("get gorm.DB failed", zap.Error(err))
		log.Fatal("gorm.db初始化失败" + err.Error())
	}

	db.SetMaxIdleConns(global.Conf.DBConfig.MaxIdleConns) // 设置空闲的最大连接数
	db.SetMaxOpenConns(global.Conf.DBConfig.MaxOpenConns) // 设置与数据库的最大打开连接数
	db.SetConnMaxLifetime(time.Hour)                      // 设置可以重用连接的最长时间
	return gormDB
}

// 初始gorm日志
func getGormLogger() logger.Interface {
	var logMode logger.LogLevel

	switch global.Conf.DBConfig.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
		LogLevel:                  logMode,                // 日志级别
		IgnoreRecordNotFoundError: false,                  // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  false,                  // 禁用彩色打印
	})
}

// 自定义 gorm Writer
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	// 是否启用日志文件
	if global.Conf.DBConfig.EnableFileLogWriter {
		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename:   global.BasePath + global.Conf.LogConfig.LogDir + "/" + global.Conf.DBConfig.LogFilename,
			MaxSize:    global.Conf.LogConfig.MaxSize,
			MaxBackups: global.Conf.LogConfig.MaxBackups,
			MaxAge:     global.Conf.LogConfig.MaxAge,
		}
	} else {
		// 默认 Writer
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}
