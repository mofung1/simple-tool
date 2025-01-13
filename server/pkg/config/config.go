package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"simple-tool/server/internal/global"
)

// Init 初始化配置
// 配置文件加载优先级：
// 1. 环境变量 CONFIG_FILE 指定的配置文件
// 2. 当前目录下的 config.yaml
// 3. 当前目录下的 config/config.yaml
// 4. 程序所在目录的 config.yaml
// 5. 程序所在目录的 config/config.yaml
func Init() {
	var err error
	v := viper.New()

	// 获取程序基础路径
	_, filename, _, _ := runtime.Caller(0)
	global.BasePath = filepath.Dir(filepath.Dir(filepath.Dir(filename)))

	// 设置配置文件类型
	v.SetConfigType("yaml")

	// 定义搜索路径
	searchPaths := []string{
		"./config.yaml",                                      // 当前目录
		"./config/config.yaml",                               // 当前目录的config子目录
		filepath.Join(global.BasePath, "config.yaml"),        // 程序所在目录
		filepath.Join(global.BasePath, "config/config.yaml"), // 程序所在目录的config子目录
	}

	// 如果环境变量中指定了配置文件，将其添加为最高优先级
	if configFile := os.Getenv("CONFIG_FILE"); configFile != "" {
		searchPaths = append([]string{configFile}, searchPaths...)
	}

	// 遍历所有可能的配置文件路径
	configLoaded := false
	for _, configPath := range searchPaths {
		v.SetConfigFile(configPath)
		if err = v.ReadInConfig(); err == nil {
			log.Printf("使用配置文件: %s\n", configPath)
			configLoaded = true
			break
		}
	}

	if !configLoaded {
		log.Printf("警告: 无法找到配置文件，将使用默认配置\n")
		return
	}

	// 把读取到的配置信息反序列化到Conf变量中
	if err := v.Unmarshal(&global.Conf); err != nil {
		log.Printf("警告: 配置文件解析失败: %v\n", err)
		return
	}

	// 监控配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Println("配置文件已更新")
		if err := v.Unmarshal(&global.Conf); err != nil {
			log.Printf("警告: 更新配置失败: %v\n", err)
		}
	})
}
