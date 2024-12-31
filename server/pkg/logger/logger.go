package logger

import (
	"log"
	"os"
	"simple-tool/server/internal/global"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Init 初始化Logger
func Init() *zap.Logger {
	writeSyncer := getLogWriter(
		global.BasePath+global.Conf.LogConfig.LogDir+"/"+global.Conf.LogConfig.Filename,
		global.Conf.LogConfig.MaxSize,
		global.Conf.LogConfig.MaxBackups,
		global.Conf.LogConfig.MaxAge,
	)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(global.Conf.LogConfig.Level))
	if err != nil {
		log.Fatal("zap.UnmarshalText() failed,err:" + err.Error())
	}

	var core zapcore.Core
	if global.Conf.Mode == "dev" {
		// 进入开发模式
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			// 日志写入文件
			zapcore.NewCore(encoder, writeSyncer, l),
			// 日志输出到终端
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		// 日志写入文件
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}

	return zap.New(core, zap.AddCaller())
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
