package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
	"net/http"
	"simple-tool/server/internal/global"
	"time"
)

type AIService struct{}

// ChatError OpenAI流式错误信息结构
type ChatError struct {
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Code    string `json:"code"`
	} `json:"error"`
}

const (
	maxMessageLength = 2000             // 最大消息长度
	streamTimeout    = 30 * time.Second // 流式请求超时时间
)

// WriteError 写入标准化的错误信息
func (s *AIService) WriteError(writer io.Writer, message, errorType, code string) error {
	chatError := ChatError{}
	chatError.Error.Message = message
	chatError.Error.Type = errorType
	chatError.Error.Code = code

	errorJSON, err := json.Marshal(chatError)
	if err != nil {
		return fmt.Errorf("错误信息序列化失败: %v", err)
	}

	_, err = fmt.Fprintf(writer, "data: %s\n\n", errorJSON)
	if err != nil {
		return fmt.Errorf("写入错误信息失败: %v", err)
	}
	if f, ok := writer.(http.Flusher); ok {
		f.Flush()
	}
	return nil
}

// validateMessage 验证消息内容
func (s *AIService) validateMessage(message string) error {
	if len(message) == 0 {
		return fmt.Errorf("消息内容不能为空")
	}
	if len(message) > maxMessageLength {
		return fmt.Errorf("消息长度超过限制")
	}
	return nil
}

// GenerateCouplet 生成对联
// name: 用户姓名
// style: 对联风格
func (s *AIService) GenerateCouplet(ctx context.Context, name, style string) (string, error) {
	config := openai.DefaultConfig(global.Conf.OpenAIConfig.APIKey)
	if global.Conf.OpenAIConfig.BaseURL != "" {
		config.BaseURL = global.Conf.OpenAIConfig.BaseURL
		global.ZapLog.Info(fmt.Sprintf("使用自定义BaseURL: %s", config.BaseURL))
	}

	client := openai.NewClientWithConfig(config)

	prompt := fmt.Sprintf("请以春节为主题，根据用户姓名'%s'和风格'%s'创作一副对联。对联要紧扣用户姓名的含义，体现新年喜庆祥和的氛围。请按照以下格式返回：\n上联\n下联\n横批", name, style)

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: "gpt-3.5-turbo",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		global.ZapLog.Error(fmt.Sprintf("调用OpenAI生成对联失败: %v", err))
		return "", fmt.Errorf("OpenAI调用失败: %v", err)
	}

	if len(resp.Choices) == 0 {
		global.ZapLog.Error("OpenAI返回结果为空")
		return "", fmt.Errorf("生成结果为空")
	}

	return resp.Choices[0].Message.Content, nil
}

// StreamChat 流式聊天
// message: 用户输入的消息
// writer: 用于写入流式响应的writer
func (s *AIService) StreamChat(ctx context.Context, message string, writer io.Writer) error {
	// 验证消息内容
	if err := s.validateMessage(message); err != nil {
		return s.WriteError(writer, err.Error(), "invalid_request_error", "invalid_message")
	}

	// 创建带超时的上下文
	timeoutCtx, cancel := context.WithTimeout(ctx, streamTimeout)
	defer cancel()

	// 创建OpenAI客户端
	config := openai.DefaultConfig(global.Conf.OpenAIConfig.APIKey)
	if global.Conf.OpenAIConfig.BaseURL != "" {
		config.BaseURL = global.Conf.OpenAIConfig.BaseURL
		global.ZapLog.Info(fmt.Sprintf("使用自定义BaseURL: %s", config.BaseURL))
	}
	client := openai.NewClientWithConfig(config)

	// 创建聊天请求
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: message,
			},
		},
		Stream: true,
	}

	// 创建流式聊天
	stream, err := client.CreateChatCompletionStream(timeoutCtx, req)
	if err != nil {
		global.ZapLog.Error(fmt.Sprintf("创建聊天流失败: %v", err))
		return s.WriteError(writer, fmt.Sprintf("创建聊天流失败: %v", err), "stream_error", "stream_create_failed")
	}
	defer stream.Close()

	// 处理流式响应
	for {
		select {
		case <-timeoutCtx.Done():
			return s.WriteError(writer, "请求超时", "timeout_error", "request_timeout")
		case <-ctx.Done():
			return s.WriteError(writer, "客户端已断开连接", "connection_error", "client_disconnected")
		default:
			response, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					// 发送结束标记
					_, err = fmt.Fprintf(writer, "data: [DONE]\n\n")
					if err != nil {
						global.ZapLog.Error(fmt.Sprintf("写入结束标记失败: %v", err))
					}
					if f, ok := writer.(http.Flusher); ok {
						f.Flush()
					}
					return nil
				}
				global.ZapLog.Error(fmt.Sprintf("接收流数据失败: %v", err))
				return s.WriteError(writer, fmt.Sprintf("接收流数据失败: %v", err), "stream_error", "stream_receive_failed")
			}

			// 检查响应的完整性
			if len(response.Choices) == 0 {
				global.ZapLog.Error("接收到的响应数据不完整")
				return s.WriteError(writer, "接收到的响应数据不完整", "stream_error", "invalid_response")
			}

			// 写入SSE格式的数据
			content := response.Choices[0].Delta.Content
			if content != "" {
				_, err = fmt.Fprintf(writer, "data: %s\n\n", content)
				global.ZapLog.Info("接收数据----" + content)
				if err != nil {
					global.ZapLog.Error(fmt.Sprintf("写入响应数据失败: %v", err))
					return s.WriteError(writer, fmt.Sprintf("写入响应数据失败: %v", err), "stream_error", "stream_write_failed")
				}
				if f, ok := writer.(http.Flusher); ok {
					f.Flush()
				}
			}
		}
	}
}
