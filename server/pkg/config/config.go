package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"path"
	"runtime"
	"simple-tool/server/internal/global"
)

// Init 初始化viper
func Init() {
	_, filename, _, _ := runtime.Caller(0)
	global.BasePath = path.Dir(path.Dir(path.Dir(filename)))

	viper.SetConfigFile(global.BasePath + "/config/config.yaml") // 指定配置文件路径
	err := viper.ReadInConfig()                                  // 读取配置信息
	if err != nil {                                              // 读取配置信息失败
		log.Fatal("viper.ReadInConfig() failed,err:" + err.Error())
	}

	// 把读取到的配置信息反序列化到Conf 变量中
	if err := viper.Unmarshal(&global.Conf); err != nil {
		log.Fatal("viper.Unmarshal failed, err:" + err.Error())
	}

	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
		if err := viper.Unmarshal(&global.Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
}
