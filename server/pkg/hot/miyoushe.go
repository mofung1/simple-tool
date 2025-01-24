package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	miyousheHotAPI = "https://bbs-api.miyoushe.com/post/wapi/getForumPostList?forum_id=1&is_good=false&is_hot=true&page_size=20&sort_type=1"
)

// MiyousheProvider 米游社热门数据提供者
type MiyousheProvider struct{}

// NewMiyousheProvider 创建新的米游社热门提供者
func NewMiyousheProvider() *MiyousheProvider {
	return &MiyousheProvider{}
}

// miyousheResponse 米游社接口响应结构
type miyousheResponse struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			Post struct {
				PostID      string `json:"post_id"`      // 帖子ID
				Subject     string `json:"subject"`      // 标题
				Content     string `json:"content"`      // 内容
				ViewNum     int    `json:"view_num"`     // 浏览数
				ReplyNum    int    `json:"reply_num"`    // 回复数
				LikeNum     int    `json:"like_num"`     // 点赞数
				ForumName   string `json:"forum_name"`   // 版块名称
				CreateTime  int64  `json:"created_at"`   // 创建时间
				Images      []string `json:"images"`     // 图片列表
				UserNickname string `json:"user_nickname"` // 用户昵称
			} `json:"post"`
		} `json:"list"`
	} `json:"data"`
}

// GetHotData 获取米游社热门数据
func (m *MiyousheProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", miyousheHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://bbs.miyoushe.com/")
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
	var miyousheResp miyousheResponse
	if err := json.Unmarshal(body, &miyousheResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 检查响应状态
	if miyousheResp.Retcode != 0 {
		return nil, fmt.Errorf("米游社API返回错误: %s", miyousheResp.Message)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range miyousheResp.Data.List {
		// 构建描述信息
		desc := fmt.Sprintf("[%s] by %s", item.Post.ForumName, item.Post.UserNickname)
		if len(item.Post.Content) > 0 {
			if len(item.Post.Content) > 100 {
				desc += "\n" + item.Post.Content[:100] + "..."
			} else {
				desc += "\n" + item.Post.Content
			}
		}
		
		result = append(result, HotData{
			Title:    item.Post.Subject,
			URL:      fmt.Sprintf("https://bbs.miyoushe.com/ys/article/%s", item.Post.PostID),
			Hot:      fmt.Sprintf("浏览%d 回复%d 点赞%d", item.Post.ViewNum, item.Post.ReplyNum, item.Post.LikeNum),
			Desc:     desc,
			Index:    i + 1,
			Platform: m.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (m *MiyousheProvider) GetPlatformName() string {
	return "米游社"
}
