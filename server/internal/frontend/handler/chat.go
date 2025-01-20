package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-tool/server/internal/frontend/service"
)

// Chat 聊天处理器
type Chat struct{}

// ChatRequest 聊天请求参数
type ChatRequest struct {
	Message string `json:"message" binding:"required"`
}

// StreamChat 流式聊天
// @Summary 流式聊天接口
// @Description 与AI进行流式聊天对话
// @Tags AI聊天
// @Accept json
// @Produce text/event-stream
// @Param message body ChatRequest true "聊天消息"
// @Success 200 {string} string "聊天响应流"
// @Router /api/v1/ai/chat [post]
func (h *Chat) StreamChat(c *gin.Context) {
	// 检查客户端是否支持SSE
	if !isSSESupported(c.Request) {
		chatService := new(service.AIService)
		_ = chatService.WriteError(c.Writer, "客户端不支持SSE", "client_error", "sse_not_supported")
		return
	}

	// 设置SSE响应头
	setSSEHeaders(c)

	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		chatService := new(service.AIService)
		_ = chatService.WriteError(c.Writer, "无效的请求参数", "invalid_request_error", "invalid_parameters")
		return
	}

	// 创建service实例并调用流式聊天方法
	chatService := new(service.AIService)
	if err := chatService.StreamChat(c.Request.Context(), req.Message, c.Writer); err != nil {
		// 错误已经在service层处理，这里不需要额外处理
		return
	}
}

// setSSEHeaders 设置SSE相关的响应头
func setSSEHeaders(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Transfer-Encoding", "chunked")
	c.Header("X-Accel-Buffering", "no") // 禁用 Nginx 缓冲
}

// isSSESupported 检查客户端是否支持SSE
func isSSESupported(r *http.Request) bool {
	accept := r.Header.Get("Accept")
	return accept == "" || // 如果没有Accept头，我们假设支持
		accept == "*/*" || // 接受所有类型
		accept == "text/event-stream" // 明确支持SSE
}
