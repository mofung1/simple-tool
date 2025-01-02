package handler

import (
	"github.com/gin-gonic/gin"
	"simple-tool/server/internal/global/response"
	"simple-tool/server/pkg/parser"
)

type Parse struct{}

func (p *Parse) Url(c *gin.Context) {
	paramUrl := c.Query("url")
	parseRes, err := parser.ParseVideoShareUrlByRegexp(paramUrl)
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}
	response.OkWithData(c, parseRes)
	return
}
