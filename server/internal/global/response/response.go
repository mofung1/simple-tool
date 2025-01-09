package response

import (
	"net/http"
	"simple-tool/server/internal/global"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var ReturnEmptyData = make(map[string]interface{}, 0)

// Result 返回结果
func Result(c *gin.Context, code int, msg string, data interface{}) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

// Ok 返回成功
func Ok(c *gin.Context) {
	Result(c, global.Success, global.Msg(global.Success), ReturnEmptyData)
}

// OkWithData 返回成功 具体数据
func OkWithData(c *gin.Context, data interface{}) {
	Result(c, global.Success, global.Msg(global.Success), data)
}

// OkWithMsg 返回成功 具体msg
func OkWithMsg(c *gin.Context, msg string) {
	Result(c, global.Success, msg, ReturnEmptyData)
}

// OkWithInfo 返回成功 具体msg及data
func OkWithInfo(c *gin.Context, msg string, data interface{}) {
	Result(c, global.Success, msg, data)
}

// OkByCode 返回错误 根据code返回
func OkByCode(c *gin.Context, code int) {
	Result(c, code, global.Msg(code), ReturnEmptyData)
}

// Fail 返回错误
func Fail(c *gin.Context) {
	Result(c, global.Error, global.Msg(global.Error), ReturnEmptyData)
}

// FailWithMsg 返回错误 具体msg
func FailWithMsg(c *gin.Context, msg string) {
	Result(c, global.Error, msg, ReturnEmptyData)
}

// FailWithInfo 返回错误 具体msg及data
func FailWithInfo(c *gin.Context, msg string, data interface{}) {
	Result(c, global.Error, msg, data)
}

// FailWithCode 返回错误 根据code返回
func FailWithCode(c *gin.Context, code int) {
	Result(c, code, global.Msg(code), ReturnEmptyData)
}
