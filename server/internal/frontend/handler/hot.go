package handler

import (
	"github.com/gin-gonic/gin"
	"simple-tool/server/internal/global/response"
	"simple-tool/server/pkg/hot"
)

// Hot 热门
type Hot struct{}

func (h *Hot) Lists(c *gin.Context) {
	paramType := c.Query("type")
	if paramType == "" {
		response.FailWithMsg(c, paramType)
		return
	}

	// 获取单个平台的数据
	provider, err := hot.NewHotDataProvider(hot.Platform(paramType))
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}
	data, err := provider.GetHotData()
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}

	//// 获取所有平台的数据
	//allData, err := hot.GetAllHotData()
	//if err != nil {
	//	// 处理错误
	//}

	response.OkWithData(c, data)
}
