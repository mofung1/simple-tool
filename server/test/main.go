package main

import (
	"fmt"
	_ "server/bootstrap"
	"simple-tool/server/internal/global"
	"simple-tool/server/internal/models"
)

func main() {
	// 自动迁移表结构
	err := global.DB.AutoMigrate(&models.Message{})
	if err != nil {
		fmt.Println("创建失败")
	}
	fmt.Println("创建完成")
}
