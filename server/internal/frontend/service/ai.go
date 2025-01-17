package service

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"simple-tool/server/internal/global"
)

type AIService struct{}

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
