package global

const (
	Success      = 200
	Error        = 400
	NotFound     = 404
	InvalidParam = 300
	ServerBusy   = 500
)

var codeMsgMap = map[int]string{
	Success:      "操作成功",
	Error:        "操作失败",
	InvalidParam: "请求参数错误",
	ServerBusy:   "服务繁忙",
	NotFound:     "请求404",
}

// Msg 状态码描述
func Msg(code int) string {
	msg, ok := codeMsgMap[code]
	if !ok {
		msg = codeMsgMap[ServerBusy]
	}
	return msg
}
