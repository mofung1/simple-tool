package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	genshinHotAPI = "https://bbs-api.miyoushe.com/post/wapi/getForumPostList?forum_id=26&is_good=false&is_hot=true&page_size=20&sort_type=1"
)

// GenshinProvider 原神热门数据提供者
type GenshinProvider struct{}

// NewGenshinProvider 创建新的原神热门提供者
func NewGenshinProvider() *GenshinProvider {
	return &GenshinProvider{}
}

// genshinResponse 原神接口响应结构
type genshinResponse struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			Post struct {
				PostID      string   `json:"post_id"`      // 帖子ID
				Subject     string   `json:"subject"`      // 标题
				Content     string   `json:"content"`      // 内容
				ViewNum     int      `json:"view_num"`     // 浏览数
				ReplyNum    int      `json:"reply_num"`    // 回复数
				LikeNum     int      `json:"like_num"`     // 点赞数
				Collection  int      `json:"collection"`   // 收藏数
				Images      []string `json:"images"`       // 图片列表
				Topics     []struct {
					Name string `json:"name"`              // 话题名称
				} `json:"topics"`
				UserNickname string `json:"user_nickname"` // 用户昵称
				CreateTime   int64  `json:"created_at"`    // 创建时间
			} `json:"post"`
		} `json:"list"`
	} `json:"data"`
}

// GetHotData 获取原神热门数据
func (g *GenshinProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", genshinHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://bbs.mihoyo.com/ys/")
	req.Header.Set("x-rpc-client_type", "4")
	
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()
	
	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}
	
	// 解析响应数据
	var genshinResp genshinResponse
	if err := json.Unmarshal(body, &genshinResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 检查响应状态
	if genshinResp.Retcode != 0 {
		return nil, fmt.Errorf("原神API返回错误: %s", genshinResp.Message)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range genshinResp.Data.List {
		// 构建话题标签
		var topics []string
		for _, topic := range item.Post.Topics {
			topics = append(topics, topic.Name)
		}
		topicStr := ""
		if len(topics) > 0 {
			topicStr = fmt.Sprintf("[%s] ", topics[0])
		}
		
		// 构建描述信息
		desc := fmt.Sprintf("%sby %s", topicStr, item.Post.UserNickname)
		if len(item.Post.Content) > 0 {
			if len(item.Post.Content) > 100 {
				desc += "\n" + item.Post.Content[:100] + "..."
			} else {
				desc += "\n" + item.Post.Content
			}
		}
		
		result = append(result, HotData{
			Title:    item.Post.Subject,
			URL:      fmt.Sprintf("https://bbs.mihoyo.com/ys/article/%s", item.Post.PostID),
			Hot:      fmt.Sprintf("浏览%d 回复%d 点赞%d 收藏%d", item.Post.ViewNum, item.Post.ReplyNum, item.Post.LikeNum, item.Post.Collection),
			Desc:     desc,
			Index:    i + 1,
			Platform: g.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (g *GenshinProvider) GetPlatformName() string {
	return "原神"
}
