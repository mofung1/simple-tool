package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"simple-tool/server/internal/global"
)

// Init 初始化配置
// 配置文件查找顺序：
// 1. 环境变量 CONFIG_FILE 指定的路径
// 2. 查找当前目录或其父级目录中的 server/config/config.yaml
func Init() {
	v := viper.New()
	var configFile string

	// 1. 首先检查环境变量
	if envConfig := os.Getenv("CONFIG_FILE"); envConfig != "" {
		configFile = envConfig
	} else {
		// 2. 获取当前工作目录
		workDir, err := os.Getwd()
		if err != nil {
			log.Fatal("无法获取当前工作目录: ", err)
		}

		// 查找包含 server/config/config.yaml 的目录
		dir := workDir
		for {
			// 尝试查找 server/config/config.yaml
			path := filepath.Join(dir, "server", "config", "config.yaml")
			if _, err := os.Stat(path); err == nil {
				configFile = path
				break
			}

			// 获取父目录
			parent := filepath.Dir(dir)
			// 如果已经到达根目录，则退出循环
			if parent == dir {
				break
			}
			dir = parent
		}

		// 如果还没找到，尝试使用可执行文件路径
		if configFile == "" {
			executable, err := os.Executable()
			if err != nil {
				log.Fatal("无法获取可执行文件路径: ", err)
			}
			executableDir := filepath.Dir(executable)

			// 尝试在可执行文件的相对位置查找配置文件
			path := filepath.Join(executableDir, "..", "config", "config.yaml")
			if _, err := os.Stat(path); err == nil {
				configFile = path
			}
		}
	}

	if configFile == "" {
		log.Fatal("未找到配置文件，请确保配置文件存在于以下位置之一：\n" +
			"1. 环境变量 CONFIG_FILE 指定的路径\n" +
			"2. server/config/config.yaml（在当前目录或其父级目录中）\n" +
			"3. ../config/config.yaml（相对于可执行文件目录）")
	} else {
		log.Println("使用配置文件-" + configFile)
	}

	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("读取配置文件失败: ", err)
	}

	// 把读取到的配置信息反序列化到Conf变量中
	if err := v.Unmarshal(&global.Conf); err != nil {
		log.Fatal("解析配置文件失败: ", err)
	}

	// 监控配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已更新")
		if err := v.Unmarshal(&global.Conf); err != nil {
			fmt.Printf("更新配置失败: %v\n", err)
		}
	})

	// 设置基础路径为 server 目录
	global.BasePath = filepath.Dir(filepath.Dir(configFile))
}
