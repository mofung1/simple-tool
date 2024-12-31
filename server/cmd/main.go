package main

import (
	"fmt"
	"net/http"
	_ "simple-tool/server/bootstrap"
	"simple-tool/server/internal/global"
	"simple-tool/server/internal/routes"
)

func main() {
	// 注册路由
	r := routes.SetupRouter()
	// 启动服务
	port := fmt.Sprintf(":%d", global.Conf.App.Port)
	if err := http.ListenAndServe(port, r); err != nil {
		global.ZapLog.Error(err.Error())
	}
}
